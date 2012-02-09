# stdinweb

I have a lot of small web services that are infrequently hit.  This
lets me run them without dedicating a process and listener to them --
I just plop them in inetd/launchd and move on.

See the [example](/dustin/whatsnew/blob/master/example/example.go) for
making your own.
