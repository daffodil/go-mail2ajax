go-mail2ajax
============

This project is experimental, under dev and not to be taken seriously, yet!

  golang newbie wants help! 
  
This app is being actively developed for use on a dedicated mailserver machine..

The Coporate App is a PyQt desktop and mobile interface, so this app
can be called from the LAN mainly

The enviroment uses deparrtmental shared mailboxes and a multi tasking
enviroment, so the system is quite "open to staff on Lan"


The app is designed to run on mailserver, and using 
- postfix
- dovecot
- postfixadmin

==================
Wishful routing..

/domains
/domain/foo.bar
/domain/foo.bar/all
/domain/foo.bar/mailbox < summary
/domain/foo.bar/mailbox/read
/domain/foo.bar/mailbox/vacation < out of office stuff
/domain/foo.bar/aliases
/domain/foo.bar/alias/<alias>