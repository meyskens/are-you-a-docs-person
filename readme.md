Are you a docs ~person~ repo?
============================

This is a small project inspired by an idea of [Carolyn Van Slyck](https://github.com/carolynvs) to calculate the amount of documentation ~contributions~ vs documentation ~contributions you have done on GitHub~ per repository.
This project makes use of [gitbase](https://github.com/src-d/gitbase) to query the repositories and [Babelfish](https://github.com/bblfsh/bblfshd) to filter any in code comments (eg. GoDoc).

## Why the pivot?
The goal of this project is to test the above listed projects, since gitbase lacked needed features I changed to working on a repo basis only looking at the HEAD.

## Results?
I did not have enough query time and resources to do everything I wanted but one metric worked fine:
```
https---github.com-BurntSushi-toml is 11.111111 % markdown
https---github.com-Masterminds-semver is 14.285714 % markdown
https---github.com-alecthomas-kingpin is 2.564103 % markdown
https---github.com-agext-levenshtein is 9.090909 % markdown
```