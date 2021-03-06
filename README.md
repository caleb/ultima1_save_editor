Ultima 1 Saved Game Editor
==========================

I recently purchased the Ultima series from [gog](http://gog.com) and decided to
play them from the beginning.

Ultima 1 is... different. You constantly need to grind for health and after
getting lost on level 5 of a dungeon and losing 5000 health and dying, I decided
that rather than grind again to beef up, I would edit the game's save files.

[Download for Mac OS X](https://dl.dropboxusercontent.com/u/473648/ultima1_save_editor-mac-os-x-2014-02-20.zip)

[Download for Windows](https://dl.dropboxusercontent.com/u/473648/ultima1_save_editor_windows.zip)

Usage
-----

    ultima1_save_editor -o PLAYER2.U1 PLAYER1.U1

The above command will open the `PLAYER1.U1` file and put the edited version in
`PLAYER2.U1`.

The program will ask you to adjust your name, stats, hit points, food, gold,
etc.

Recommendations
---------------

I haven't tested the exact limits, but I wouldn't assign more than 200 or so for
your stats (strength, agility, etc).

You can give yourself 999 intelligence, but this causes items in shops to be
incredibly expensive since shop prices are based on intelligence, and having
such a high int much cause an overflow. A 200 int is enough to drop those prices
to 0.

Save files
----------

On Mac OS X, the save files for the GOG version of Ultima 1 are located in:

    ~/Library/Application Support/Boxer/Gamebox States/com.gog.ultima-1/Current.boxerstate/C Ultima1.harddisk

I'm guessing that this program will work with any save file generated by the
Ultima 1 remake from 1986
