# High level overview of Nebula - interchain launchpad

Firstly, this is for my graduation thesis.

I observe Cosmos and realize that there is no true launchpad in Cosmos.

There are following problems a launchpad needs to solve:
1. distribution of exclusive position

for each projects, exclusive position is limited, a mechanism to distribute these exclusive positions is needed so as to be as fair as possible.

2. filter bots

because entering a project in the beginning is very lucrative, a lot will try to cheat the system with fake accounts so as to gain as many benefits as possible.

3. many ways to launch a project

this is what special about this interchain launchpad. An interchain launchpad should have as many ways to call funds for a project as possible.

# Composition
1. alloc: how to mint and distribute tokens.
2. claim: claim rewards for airdrop.
3. launchpad: a module that governs projects on interchain launchpad.
4. distribution mechanism (each will require a new module): various distribution mechanisms to distribute project token to early users:
* ido
* lock_drop
5. participate: a module that governs user interaction towards a launchpad project.