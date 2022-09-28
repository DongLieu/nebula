# Release Mechanism Interface

As mentioned in [launchpad](/x/launchpad/README.md), release mechanism is an interface required by a launchpad project.

For maximum flexibility, release mechanism interface should strive to describe as broad of a release mechanism as possible.

In broad term, release mechanism is just a way to release assets of varied attributes and varied quantities to user under a specific mechanism.

Because, each of these release mechanism will have different ways for user to interact with it. Therefore, each release mechanism must include user interaction tx as well.

# Characteristic
1. Flexibility: release mechanism is designed with flexibility in mind so as it can upgrade later.
1. Secure: release mechanism must ensure security for investor and project assets.

# Structure overview
1. asset: there can be many asset types in the future. For now, it will be token. NFT will be supported later.
    * In case of NFT, I need project address to store NFTs. I should probably expose project address as well.
1. total_amount: total amount of assets for release.
1. total_released: amount of tokens already released.
1. asset_listing_price: an asset needs to have listing price on launchpad in US dollar.
1. asset_accepted_payment: an asset needs to declare what types of asset it accepts as means of payment. 
    * The only accepted token is stable coin.
1. asset_released_per_entry: limit on how much assets will be released for an entry.
1. end_condition: end condition for release mechanism
    * Project ends when __total_amount__ - __amount_released__ = 0.
    * Project ends when time ends
    * Project reaches a fundraising goal
1. commitment_criteria: criteria for an entry to be eligible
1. participants: people that have participated

# ReleaseMechanism interface that allows project to interact with
1. DeleteReleaseMechanism
2. GetTokens
3. GetReleaseMechanismStatus
    * 0: has not started
    * 1: active
    * 2: ended

# Implementation
1. [IDO](/x/ido/README.md) 