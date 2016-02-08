/*
 * MumbleDJ
 * By Matthieu Grieger
 * strings.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

// Current version of the bot 
const VERSION = "08/15"

// Message shown to users when they request the version of the bot 
const DJ_VERSION = "Rotomat <b>" + VERSION + "</b>"

// Message shown to users when the bot has an invalid API key.
const INVALID_API_KEY = "Rotomat hat keinen gültigen %s API-Key."

// Message shown to users when they do not have permission to execute a command.
const NO_PERMISSION_MSG = "Du hast keine Rechte für diesen Befehl."

// Message shown to users when they try to add a playlist to the queue and do not have permission to do so.
const NO_PLAYLIST_PERMISSION_MSG = "Du hast keine Rechte für das Hinzufügen von Playlists."

// Message shown to users when they try to execute a command that doesn't exist.
const COMMAND_DOESNT_EXIST_MSG = "Der eingegebene Befehl existiert nicht."

// Message shown to users when they try to move the bot to a non-existant channel.
const CHANNEL_DOES_NOT_EXIST_MSG = "Der angegebene Kanal existiert nicht."

// Message shown to users when they attempt to add an invalid URL to the queue.
const INVALID_URL_MSG = "Der Link den du eingesendet hast, ist in keinem unterstützten Format."

// Message shown to users when they attempt to search on an invalid platform.
const INVALID_SEARCH_PROVIDER  = "The Search provider you submitted does not match the required format."

// Message shown to users when they attempt to add a video that's too long
const TRACK_TOO_LONG_MSG = "Deine Einsendung %s überschreitet die maximal festgelegte Länge."

// Message shown to users when they attempt to perform an action on a song when
// no song is playing.
const NO_MUSIC_PLAYING_MSG = "Aktuell wird keine Musik abgespielt."

// Message shown to users when they attempt to skip a playlist when there is no playlist playing.
const NO_PLAYLIST_PLAYING_MSG = "Aktuell wird keine Playlist abgespielt."

// Message shown to users when they try to play a playlist from a source which doesn't support playlists.
const NO_PLAYLISTS_SUPPORTED_MSG = "Playlisten von %s werden nicht unterstützt."

// Message shown to users when they attempt to use the nextsong command when there is no song coming up.
const NO_SONG_NEXT_MSG = "Aktuell sind keine Titel in der Warteschlange."

// Message shown to users when they issue a command that requires an argument and one was not supplied.
const NO_ARGUMENT_MSG = "Der eingebene Befehl benötigt mehr Angaben, die du weggelassen hast."

// Message shown to users when they try to change the volume to a value outside the volume range.
const NOT_IN_VOLUME_RANGE_MSG = "Bitte was? Die Lautstärke muss zwischen %f und %f sein."

// Message shown to user when a successful configuration reload finishes.
const CONFIG_RELOAD_SUCCESS_MSG = "Alle Einstellungen neu geladen."

// Message shown to users when an admin skips a song.
const ADMIN_SONG_SKIP_MSG = "Ein Admin hat den Titel übersprungen."

// Message shown to users when an admin skips a playlist.
const ADMIN_PLAYLIST_SKIP_MSG = "Ein Admin hat die Playlist übersprungen."

// Message shown to users when the audio for a video could not be downloaded.
const AUDIO_FAIL_MSG = "Der Audio-Download für dieses Video ist fehlgeschlagen. Springe zum nächsten Titel!"

// Message shown to users when they supply an URL that does not contain a valid ID.
const INVALID_ID_MSG = "Der eingesendete Link %s enthält keine gültige ID."

// Message shown to user when they successfully update the bot's comment.
const COMMENT_UPDATED_MSG = "Der Kommentar für den Bot wurde erfolgreich aktualisiert."

// Message shown to user when they request to see the number of songs cached on disk.
const NUM_CACHED_MSG = "Aktuell werden %d Titel auf der HDD gecached."

// Message shown to user when they request to see the total size of the cache.
const CACHE_SIZE_MSG = "Der Cache beträgt aktuell %g MB."

// Message shown to user when they attempt to issue a cache-related command when caching is not enabled.
const CACHE_NOT_ENABLED_MSG = "Der Cache ist aktuell deaktiviert."

// Message shown to user when they attempt to shuffle the queue and it has less than 2 elements.
const CANT_SHUFFLE_MSG = "Shuffeln funktioniert nicht bei weniger als 2 Titeln."

// Message shown to users when the songqueue has been successfully shuffled.
const SHUFFLE_SUCCESS_MSG = "Die aktuelle Warteschlange wurde geshuffelt durch <b>%s</b> (gültig ab dem nächsten Titel)."

// Message shown to users when automatic shuffle is activated
const SHUFFLE_ON_MESSAGE = "<b>%s</b> hat automatisches Shuffeln aktiviert."

// Message shown to users when automatic shuffle is deactivated
const SHUFFLE_OFF_MESSAGE = "<b>%s</b> hat automatisches Shuffeln deaktiviert."

// Message shown to user when they attempt to enable automatic shuffle while it's already activated
const SHUFFLE_ACTIVATED_ERROR_MESSAGE = "Automatisches Shuffeln ist bereits aktiviert."

// Message shown to user when they attempt to disable automatic shuffle while it's already deactivated
const SHUFFLE_DEACTIVATED_ERROR_MESSAGE = "Automatisches Shuffeln ist bereits deaktiviert."

// Message shown to channel when a song is added to the queue by a user.
const SONG_ADDED_HTML = `
	<b>%s</b> hat "%s" zur Warteschlange hinzugefügt.
`

// Message shown to channel when a playlist is added to the queue by a user.
const PLAYLIST_ADDED_HTML = `
	<b>%s</b> hat die Playliste "%s" zur Warteschlange hinzugefügt.
`

// Message shown to channel when a song is added to the queue by a user after the current song.
const NEXT_SONG_ADDED_HTML = `
	<b>%s</b> hat "%s" zur Warteschlange (nach dem aktuellen Titel) hinzugefügt.
`

// Message shown to channel when a playlist is added to the queue by a user after the current song.
const NEXT_PLAYLIST_ADDED_HTML = `
	<b>%s</b> hat die Playlist "%s" zur Warteschlange (nach dem aktuellen Titel) hinzugefügt.
`

// Message shown to channel when a song has been skipped.
const SONG_SKIPPED_HTML = `
	Stimme gegen den aktuellen Titel abgegeben.
`

// Message shown to channel when a playlist has been skipped.
const PLAYLIST_SKIPPED_HTML = `
	Stimme gegen die aktuelle Playlist abgegeben.
`

// Message shown to display bot commands.
const HELP_HTML = `<br/>
	<b>Benutzerbefehle:</b>
	<p><b>!help</b> - Zeigt diese Hilfe an.</p>
	<p><b>!search (yt|sc) query</b> - Search on Youtube or Soundcloud for a query and add first hit.</p>
	<p><b>!add</b> - Fügt Songs/Playlisten hinzu.</p>
	<p><b>!v</b> - Zeigt entweder die aktuelle Lautstärke an oder ändert sie.</p>
	<p><b>!skip</b> - Stimme gegen den aktuellen Titel abgeben.</p>
	<p> <b>!skipplaylist</b> - Stimme gegen die aktuelle Playlist abgeben.</p>
	<p><b>!numsongs</b> - Zeigt an, wieviele Titel in der Warteschlange sind.</p>
	<p><b>!list</b> - Zeigt die Songs in der Warteschlange an.</p>
	<p><b>!nextsong</b> - Zeigt Titel und Einsender des nächsten Songs an, wenn es einen gibt.</p>
	<p><b>!currentsong</b> - Zeigt Titel und Einsender des aktuellen Songs an.</p>
	<p><b>!version</b> - Zeigt die Version von Rotomat an.</p>
`

// Message shown to users when they ask for the current volume (volume command without argument)
const CUR_VOLUME_HTML = `
	Die Lautstärke ist <b>%.2f</b>.
`

// Message shown to users when another user votes to skip the current song.
const SKIP_ADDED_HTML = `
	<b>%s</b> hat gegen den aktuellen Titel gestimmt.
`

// Message shown to users when the submitter of a song decides to skip their song.
const SUBMITTER_SKIP_HTML = `
	Der Einsender hat den Titel übersprungen.
`

// Message shown to users when another user votes to skip the current playlist.
const PLAYLIST_SKIP_ADDED_HTML = `
	<b>%s</b> hat gegen die aktuelle Playlist gestimmt.
`

// Message shown to users when the submitter of a song decides to skip their song.
const PLAYLIST_SUBMITTER_SKIP_HTML = `
	Die aktuelle Playlist wurde durch den Einsender <b>%s</b> übersprungen.
`

// Message shown to users when they successfully change the volume.
const VOLUME_SUCCESS_HTML = `
	<b>%s</b> hat die Lautstärke auf <b>%.2f</b> geändert.
`

// Message shown to users when a user successfully resets the SongQueue.
const QUEUE_RESET_HTML = `
	<b>%s</b> hat die Warteschlange geleert.
`

// Message shown to users when a user asks how many songs are in the queue.
const NUM_SONGS_HTML = `
	Aktuell <b>%d</b> Titel in der Warteschlange.
`

// Message shown to users when they issue the nextsong command.
const NEXT_SONG_HTML = `
	Der nächste Song in der Warteschlange ist "%s", hinzugefügt durch <b>%s</b>.
`

// Message shown to users when they issue the currentsong command.
const CURRENT_SONG_HTML = `
	Der aktuelle Titel ist "%s", hinzugefügt durch <b>%s</b>.
`

// Message shown to users when the currentsong command is issued when a song from a
// playlist is playing.
const CURRENT_SONG_PLAYLIST_HTML = `
	Der aktuelle Titel ist "%s", hinzugefügt durch <b>%s</b> über die Playlist "%s".
`

// Message shown to user when the listsongs command is issued
const SONG_LIST_HTML = `
	<br>%d: "%s", hinzugefügt durch <b>%s</b>.</br>
`
