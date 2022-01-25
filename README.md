Written by Damen Knight

Copyright 2021

This is just a repo I setup to share some of the code I've written for the newest iteration of dogecointicker.  This iteration is entirely GoLang, and runs off a Windows PC, which is why a lot of it is SUPER funky.

Plans include: Updated Art generation, non-Windows code, actual setup/usage instructions.

List of files and what they do:

dogeartgenerator.go - this generates the daily artwork that the bot uploads

dogeartuploader.go - this uploads/tweets out the art that was generated

dogecoin2gum.go - this figures out how much gum you can buy with the current dogecoin value, generates the artwork, and uploads it

dogecoinfud.go - this tweets out the daily "random" FUD about dogecoin

dogecointicker.go - this is the main code for the bot.  This does all the regular value tweets

dogecointickerresponder.go - this handles replying to DMs with random fun facts about dogecoin

Current cron configuration
*/60 * * * * /<path>/dogecointicker
30 17 * * * /<path>/dogecoinfud
*/5 * * * * /<path>/dogecointickerresponder
05 12 * * * /<path>/dogecoin2gum
00 5 * * * /<path>/dogeartgenerator
30 9 * * * /<path>/dogeartuploader

