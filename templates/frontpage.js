var calendar = {};
calendar.globals = {};

calendar.parse_time = function (item) {
	var hour = 0;
	var minute = 0;
	item = item.toLowerCase();
	if (item.indexOf('am') != -1) {
		item = item.split(item.indexOf('am'))[0];
	}
	if (item.indexOf('pm') != -1) {
		item = item.split(item.indexOf('pm'))[0]
		hour += 12;
	}
	if (item.indexOf(':') != -1) {
		hour += parseInt(item.split(':')[0]) || 0;
		minute += parseInt(item.split(':')[1]) || 0;
	} else {
		hour += parseInt(item) || 0;
	}
	return hour * 60 + minute;
}

calendar.add = function () {
	var date = Date.parse($('#date').val());
	var time = calendar.parse_time($('#time').val());
	var item = $('#item').val();
	console.log(date, time, item);
	calendar.show_main();
};

calendar.hide_all = function () {
	$('#main').hide();
	$('#add').hide();
};

calendar.show_add = function () {
	calendar.hide_all();
	$('#add').show();
};

calendar.build_add = function () {
	$('#calendar').append($('<div id=add></div>'));
	$('#add').append($('<p>Date</p>'));
	$('#add').append($('<input id="date"></input>'));
	$('#add').append($('<p>Time</p>'));
	$('#add').append($('<input id="time"></input>'));
	$('#add').append($('<p>Item</p>'));
	$('#add').append($('<input id="item"></input>'));
	$('#add').append($('<input type=submit onclick=calendar.add()></input>'))
	$('#add').hide();
};

calendar.build_main = function () {
	$('#calendar').append($('<div id=main></div>'));
	$('#main').append($('<input type=submit onclick=calendar.show_add() value=add></input>'))
	$('#main').hide();
};

calendar.show_main = function () {
	calendar.hide_all();
	$('#main').show();
}

$(document).ready(function() {
	calendar.build_add();
	calendar.build_main();
	calendar.show_main();
})
