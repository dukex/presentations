var express   = require('express');
var fs        = require('fs');
var io        = require('socket.io');
var _         = require('underscore');
var Mustache  = require('mustache');

var app       = express.createServer();
var staticDir = express.static;

app.use(express.cookieParser());
app.use(express.session({secret: 'mysecretXXX'}));

io            = io.listen(app);

var opts = {
	port :      1947,
	baseDir :   __dirname + '/../../'
};

io.sockets.on('connection', function(socket) {
	socket.on('slidechanged', function(slideData) {
		socket.broadcast.emit('slidedata', slideData);
	});
	socket.on('fragmentchanged', function(fragmentData) {
		socket.broadcast.emit('fragmentdata', fragmentData);
	});
});

app.configure(function() {
	[ 'css', 'js', 'images', 'plugin', 'lib', 'videos' ].forEach(function(dir) {
		app.use('/' + dir, staticDir(opts.baseDir + dir));
	});
});


app.get("/notes/:socketId", function(req, res) {
	fs.readFile(opts.baseDir + 'plugin/notes-server/notes.html', function(err, data) {
		res.send(Mustache.to_html(data.toString(), {
			socketId : req.params.socketId,
      presentation_uri: req.session.presentation_uri
		}));
	});
});

app.get("/favicon.ico", function(req, res) {
	res.writeHead(404, {'Content-Type': 'text/html'});
})

app.get("/:presentation_uri", function(req, res) {
	res.writeHead(200, {'Content-Type': 'text/html'});
  req.session.presentation_uri = req.params.presentation_uri;
	fs.createReadStream(opts.baseDir + '/' + req.params.presentation_uri).pipe(res);
});


// Actually listen
app.listen(opts.port || null);

var brown = '\033[33m',
	green = '\033[32m',
	reset = '\033[0m';

var slidesLocation = "http://localhost" + ( opts.port ? ( ':' + opts.port ) : '' );

console.log( brown + "reveal.js - Speaker Notes" + reset );
console.log( "1. Open the slides at " + green + slidesLocation + reset );
console.log( "2. Click on the link your JS console to go to the notes page" );
console.log( "3. Advance through your slides and your notes will advance automatically" );
