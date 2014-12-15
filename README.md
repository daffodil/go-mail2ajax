go-mail2ajax
============

This project is experimental, under dev and not to be taken seriously.

Quick Rundown..

This is an experimental idea to create an interface
to both admin a mailserver,
but also to send emails and read via 
a AJAX REST interface, all in one GoLang app 

So the idea is:
- a go http server spits out ajax
- a imap connector to read mailboxes
- a way to change profile, eg passwd, out of office, forwardings
- a rest connector to CRUD mailboxes

This app is currently under development for use
at a couple of small businesses with lots of email and lots
and lots of attachments flying around.

The service is intended fo use via json with pyqt desktop application
an android/phonegap app and a jquerymobile app.

Help wanted as pedro is a golang newbie.


