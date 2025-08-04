# Contest Platform

This document will serve as a design document for a contest platform.

# Overview

By the end of said project, it should be a fully functioning cp contest platform that displays the problems, gets the contestants responses, runs them, returns verdicts, all while providing a leaderboard in real time.
It should also support multiple languages, and it goes without saying that the contest problems are bound by memory and time constraints.

for the time being, we will concentrate on a python venv, then work our way from there ...

# Design analysis

## Overall architecture
Multiple language support makes a microservice approach more likely: basically a layout where a dispatcher entity takes the contestants entries and serves as a work cache for each language worker to take over.

There will also be another supervising entity *Supervisor*, its role is to start a dispatch entity, cache contestants info, and manage the leaderboard.

Now that the relatively obvious part is cleared, onto the workers themselves. Time and memory constraints dictate that concurrency is out of question, and a separation according to languages must be in place (whether that separation should be physical or not clear for me for now, since if it is so, that means that concurrency should not be a problem, we will see about that later ... ).
