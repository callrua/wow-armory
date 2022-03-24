# wow-armory

![](https://github.com/callrua/wow-armory/workflows/Build/badge.svg)

### Problem:

World of Warcraft (PvP) is considered a "solved" game, i.e. there is almost always an optimized way to play the game/build your character. As such, I find myself constantly checking the top players to see how they are gearing.
I have also recently been exposed to Go development and was keen to supplement with a small project.

### Concept: 

I want to make use of the [World of Warcraft Game API](https://develop.battle.net/documentation/world-of-warcraft/game-data-apis) to pull info about the top X players on ladder.

Should be passed in a class, potentially a number and will output the following:

 * Their Legendary x 2 and which slot they are in.
 * Stats allocation.
 * Gems/Enchants. (?)

If possible, should be able to pass in spec too. So that in cases where one spec is better than the others we don't have to get 100 Fire Mages for every 2 Frost Mages.

Initial idea for flow would be:

1. Gather the top X players for the spec, from the leaderboard
2. Go routines to pull their armory and parse into an output.

Implementing a `BAZEL.build` and a `Dockerfile` to tinker with some build tools and a `Github Workflow` to play with some CI/CD concepts.

I'm happy with just running this via CLI for now - but if it could be hosted somewhere that'd be neat as realistically anywhere I have a terminal I won't have WoW.
