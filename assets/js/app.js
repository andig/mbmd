let sort = {
	methods: {
		sorted: function (theMap) {
			var devs = Object.keys(theMap);
			devs.sort();
			var res = {};
			devs.forEach(function (key) {
				res[key] = theMap[key];
			});
			return res;
		}
	}
}

let formatter = {
	methods: {
		// val returns addable value: null, NaN and empty are converted to 0
		val: function (v) {
			v = parseFloat(v);
			return isNaN(v) ? 0 : v;
		},
		fmt: function (v) {
			return this.val(v).toFixed(2);
		}
	}
}

let dataapp = new Vue({
	el: '#realtime',
	delimiters: ['${', '}'],
	mixins: [sort],
	data: {
		meters: {},
		message: 'Loading...'
	},
	computed: {
		sortedMeters: function() {
			console.log(this.meters)
			return this.sorted(this.meters);
		}
	}
});

Vue.component('row', {
	template: '#measurement',
	delimiters: ["${", "}"],
	mixins: [formatter],
	props: ["data", "title", "sum"],
	computed: {
		valsum: function () {
			if (this.total !== undefined) {
				return this.total;
			} else {
				return this.val(this._1) + this.val(this._2) + this.val(this._3);
			}
		},
	},
});

let statusapp = new Vue({
	el: '#status',
	delimiters: ['${', '}'],
	mixins: [sort],
	data: {
		meters: {}
	}
});

$().ready(function () {
	connectSocket();
});

function updateStatus(status) {
	var id = status["Device"]
	status["Status"] = status["Online"] ? "online" : "offline"

	// update data table
	var dict = statusapp.meters[id] || {}
	dict = Object.assign(dict, status)

	// make update reactive, see
	// https://vuejs.org/v2/guide/reactivity.html#Change-Detection-Caveats
	Vue.set(statusapp.meters, id, dict)
}

const re = /^(.+?)([SL]([1-9]))?$/

let fixed = d3.format(".2f")
let si = d3.format(".3~s")

function updateData(data) {
	// extract the last update
	let id = data["Device"]
	let type = data["IEC61850"]
	let value = fixed(data["Value"])

	// put into status line
	dataapp.message = "Received " + id + " / " + type + ": " + si(value)

	// match type
	let match = re.exec(type)
	let base = match[1]
	let component = "total"
	if (match[3] !== undefined) {
		component = "_" + match[3]
	}

	// create or update data table
	let meter = dataapp.meters[id] || {}
	let dict = meter[base] || {}

	dict[component] = value
	meter[base] = dict

	// make update reactive, see
	// https://vuejs.org/v2/guide/reactivity.html#Change-Detection-Caveats
	Vue.set(dataapp.meters, id, meter)
}

function processMessage(data) {
	if (data.Meters && data.Meters.length) {
		for (var i=0; i<data.Meters.length; i++) {
			updateStatus(data.Meters[i]);
		}
	}
	else if (data.Device) {
		updateData(data);
	}
}

function connectSocket() {
	var ws, loc = window.location;
	var protocol = loc.protocol == "https:" ? "wss:" : "ws:"

	// ws = new WebSocket(protocol + "//" + loc.hostname + (loc.port ? ":" + loc.port : "") + "/ws");
	ws = new WebSocket("ws://localhost:8081/ws");

	ws.onerror = function () {
		ws.close();
	}
	ws.onclose = function () {
		window.setTimeout(connectSocket, 1000);
	}
	ws.onmessage = function (evt) {
		// console.log(evt.data);
		var json = JSON.parse(evt.data);
		processMessage(json);
	}
}
