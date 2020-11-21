# Cruel  MIT Board

## Intro
This is a program that aims to learn online courses from [MIT OpenCourseWare](https://ocw.mit.edu/index.htm).

The learning course range may be expanded later.

The assessment is slightly cruel.

### Code Structure
- Checkin Result Board: a website that shows checkin data for each member.
- Crawler: A Github Crawler that runs periodically for checkins.

### Some Terms
- `Small Red Pocket`: `floor(10*e)`(27) RMB, distributed to 15 red pockets randomly.
- `Big Red Pocket`: `floor(20*e)`(54) RMB, distributed to 30 red pockets randomly.
- `peer`: It takes the round-robin strategy. A peer of a member refers to the one previous to him/her in the WeChat Group. The peer of the group owner is the latest member of the group, so the work of the latest member will be checked by the owner.


## Join/Leave the Group

### Join
1. A WeChat group is created, and the capacity is fixed at `floor(10*e)`(27).

2. One can join the group as long as the group is not full, having finished the `prerequisites for join group` and being invited by someone in the group.

3. To join the group, **one should provide his/her GitHub id** to the group owner.

4. One can choose to make a reservation with the group owner and start from the next course. Or catch up with the current course and join the moment the requirements are met, a manual inspection request should be made to the group owner in this case.

#### Prerequisites for Join Group
1. [MIT 6.null](https://missing.csail.mit.edu/) (But if the current course is 6.null, you don't have to take this course in advance).

### Leave

A member have to pay a `big red pocket` if he/she wants to leave the group.

### Invitation

One have to finish at least one course before he/she quit the group, otherwise the one who invites him/her should offer a `small red pocket`. (The group owner is excluded from this term before the number of members of the group reaches 27 for the first time)

## Learning Repository

To make the checkin more convenient, we make a Github Learning Repository for each course mandatory.

For example, if we start learning MIT 6.null. The repository name should be `mit-6-null`(this will be specified by the group owner in advance).
 
 
The repository should be like the following.

```
week1
    mon
        README.md
    tue
        README.md
    wed
        README.md
    thur
        README.md
    fri
        README.md
    sat
        README.md
    sun
        README.md
    task
        ...(It depends on you and assignment requirements)
    peer-review
        README.md
week2
    mon
    tue
    wed
    thur
    fri
    sat
    sun
    task
    peer-review
week3
    mon
    tue
    wed
    thur
    fri
    sat
    sun
    task
    peer-review
......(If there are more weeks for this week)
```

### Diary Interpretation
A valid diary Example:
- Mon: Today I watched half of the video
- Tue: (vacant because today is busy)
- Wed: Today I watched the other half of the video
- Thur: I read the assignments
- Fri: (vacant since there's a party with friends)
- Sat: Finished the assignments
- Sun: (Peer Review and the start of the next lecture)

The diary does not require much effort, and it just makes sure that you're working on the lecture and assignments. You may also add some of your learning thoughts to it.

 
## Checkin

All the assignments or learning diary should be updated to the registered Github Account as checkin.

Meanwhile, each checkin should be accompanied by a screen shot sent to the WeChat group.

A checkin missing either one of the above is not a valid checkin. 

All the checkins MUST BE in English, anyone who violates it should offer a `smaller red pocket`.

### Checkin Detection

The creation of the new folder will be checked every day. A Group member of the group MUST checkin AT LEAST each other day. 

Each detection will be for the current day and the day before. Either one or both checkins will be valid.

### Weekly Task & `Peer` Review
The weekly task is published at the beginning of each week, which typically is an one hour long lecture and its accompanied assignments. 

Each Sunday, there will be a peer review. One have to create a report in his or her repository, commenting on his/her `peer`.

## Terms for Members

1. One should be aware that it's his/her responsibility to have learnt the prerequisites for each course, including the newly joined ones.
2. Spreading Anxiety is invalid, and one who violates this term should offer a `big red pocket`.

## Rewards and Punishment 
- A member who does not checkin in consecutive 2 days should offer a `small red pocket`.
- A member who does not finish the weekly task should offer a `big red pocket`.
- If a member is reported to be checkin badly, he/she should pay double the amount of the red pocket should have paid originally for the corresponding checkin.

## Our Mission
We have a cause for this program - Lifelong Learning.

We have a slogan for our mission - YOUXIU is our Habit.