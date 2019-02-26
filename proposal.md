# Project Proposal

> This project's direction is still uncertain at the moment.
> However, it will be a slack bot. I have been given some interesting information:
> Slack's *Slack Tank* is an idea engine where employees put out their ideas internally.
> Some of these would definately be an interesting idea to use for this go project.
> All of the data below this is quoted text!


:slacktank-2019: *PITCH TITLE:* Rightsize Your Rooms :slacktank-2019:
*Category:* company (app)
*Team Members:* @allyc @Rowen Hahn
*Slide Deck:* N
*Video:* N
*I/we plan to hack this ourselves:* N

Ever needed a conference room for 10 people, but they’re all booked by 1:1s? What if I told you there’s a bot that lives in #find-rooms-500 and you can type `/rightsize thurs 2pm 10 people` and it finds you a 10 person conference room with fewer than five SF participants! How would you feel if it asks you if you want to make a swap and then starts a DM-negotiation with the room holder and a workplace rooms-and-spaces person? You’d all get the Right Sized Rooms and we’d never worry about meeting space again!

I love this one…the bot can look up the meeting booked on GCal and figure out the participant count.
:brain: Could even get sophisticated and group participants by profile location and accurately calculate for multi-location, multi-room meetings.
:letsroll: Seems do-able if we can look up GCal meeting info through APIs.
_Here’s my vision for how to automate this process:_

• It would be *amazing* if I could simply type `/agenda` in each 1-on-1 channel and automagically get my customized agenda for that meeting.
• I am envisioning this working similarly to our Triage Bot, where it would scan the channel looking for messages that met certain criteria (specifically, messages that had reacji corresponding to defined agenda sections)
•The action would be to create a new Post (based on a copy of the “master” agenda) and insert a copy of each messages into the appropriate section of an agenda.

_Why this matters:_

This use case has incredible potential. Right now I am the only person doing this because it is so cumbersome. If we streamlined the process, I could see other managers at Slack adopting this as their process for running 1-on-1s.

More importantly, this could become a really powerful use case for our customers. I believe front-line managers are a critical linchpin to the adoption of any tool or business process, because they can activate their direct reports, their peers, and their leadership. 1-on-1 meetings are ubiquitous, so unlocking this use case could have profound implications for adoption of Slack across our customer base.

Please help me make this a reality! :t-hanks:
:slacktank-2019::microscope: *Microbot, an internal microblogging platform* :microscope::slacktank-2019:
*Category*: Product
*Team members:* @shinypb, @anuj
*Slide deck:* N
*We plan to hack this ourselves*: Y

Microbot is a bot for internal microblogging — a place to post short thoughts that anyone else on your team can follow if they want to. You'll DM your thoughts to @microbot, and you can follow any of your team mates. Messages from people that you follow send to _their_ @microbot will show up as message inside of _your_ @microbot.

Minimum viable product: just the @microbot bot.
Maximum viable product: a top-level microblog view, along the lines of All Unreads.

(Note: I've actually built something similar to this as a side project; this publicly-accessible microblog is actually built from a DM with a bot on my personal Slack team. It makes for a pretty compelling authoring experience. https://whimsicalifornia.com/microbot/user/shinypb )
slacktank-2019:*Pitch Title*: *Personal Feedback Bot* :slacktank-2019:
*Category:* Product
*Team Members:* None yet :sad_penguin:
*Slide Deck Included:* N
*Video Included:* N
*I/we plan to hack this ourselves:* Would love to be involved, but :zero: development skills!!

Have you ever walked out of a meeting/presentation/discussion with a note to self to provide feedback to a peer and forgot? Is semi-annually too long to wait to be able to provide feedback to your peers? Do you just want to let someone know they just killed it?!

Introducing the 'Personal Feedback Bot', just one small step toward
>"*Clarity is kindness* - Talk *to* people, not *about* people. Give direct feedback."

Provide feedback directly in Slack to a peer, anonymously or be known, and that feedback comes to the recipient as a DM from the Bot. Nice to have - you can access that feedback and copy it into your semi-annual self-assessment!
