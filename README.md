go-mail2ajax
============

This project is experimental, under dev and not to be taken seriously, yet!

  golang newbie wants help! 
  
The app is a revised/updated version of a realworld app, running on a
mailserver and attached database, eg a dedicated machine or a LAN+online integration.


Read Access Mailboxes: 
---------------------------

/ajax/mailbox/{my@example.com}/summary
- returns a list of folder ++ top + latest messages

/ajax/mailbox/{my@example.com}/folders
- returns a list of imap folders and new mesages etc

/ajax/mailbox/{me@address.com}/folder/{imap_folder}/message/{uid}
- returns the message embedded in json, in peek later





