# IDO module

this is intended as an implementation of ReleaseMechanism

IDO allows people to buy token of the project at pre - release price. It helps founder team to raise fund.

IDO needs a way to select people else a lot of people will buy like crazy.

## Ways to add tokens into project
1. add through personal address: this way adds tokens into the project from personal address.
    * need to check if such token amount exists in personal address
    * can be rather limiting in terms of token availability
        * new project that doesn't launch yet so no tokens on mainnet
        * a project has not enabled IBC to this launchpad

2. add through token factory: this way adds token into the project from token factory.
    * project owner can generate any amount of token into their project vault
    * high token availability
    * token generated through token factory shall not leave this launchpad through IBC in order not to confuse with the project's real token
    * there will be a list of token holders so that the real project can airdrop 1:1 to these holders
    * (feature request) the ability to convert factory token to main project token through another module "convert"

a project can either add token through way 1 or way 2. Both is not allowed.

# Project Version
1. v1: create ido
2. v2: participate in ido freely (this will be the scope for demo)
3. v3: participate in ido with select mechanism

# IDO structure
1. project_id: the project id that this IDO is linked to
1. token_for_distribution: the total amount of coins for distribution
    * keeping a local record here helps IDO executes logic directly on IDO without the need to retrieve assets information in launchpad. After logic is executed, balances in launchpad module's project will be updated accordingly
1. total_distributed_amount: total amount for distribution
1. token_listing_price: the initial price of token (in US Dollar)
1. allocation_limit: how much token worth is allowed to purchase (in US Dollar)
1. ido_status: status of IDO
1. start_time: start time of IDO
1. entries: people that joins this IDO

# start - end condition
1. start condition: when block time > start time of a project
1. end condition:
    * (if has end time) end time is reached if block time > end time of a project
    * token_for_distribution == 0, ido_status == active

# Global params
# Message Tx
1. EnableIDO: tx message to enable ido for a specific project id
1. DisableIDO: tx message to disable ido for a specific project id
	* ido status has to be 0
	* project's all release mechanism status is 0
2. CommitParticipation: tx message to buy tokens
    * token_commit: the amount of buy token to be commited
    * can only commit if the project is active

# Query
1. TotalAmountDistributed: query total amount distributed so as to know the release progress
2. IDO: query ido information corresponding to a project id, this is intended to check and get release mechanism of a project
