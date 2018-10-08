Are you a docs ~person~ repo?
============================

This is a small project inspired by an idea of [Carolyn Van Slyck](https://github.com/carolynvs) to calculate the amount of documentation ~contributions~ vs documentation ~contributions you have done on GitHub~ per repository.
This project makes use of [gitbase](https://github.com/src-d/gitbase) to query the repositories and [Babelfish](https://github.com/bblfsh/bblfshd) to filter any in code comments (eg. GoDoc).

## Why the pivot?
The goal of this project is to test the above listed projects, since gitbase lacked needed features I changed to working on a repo basis only looking at the HEAD.