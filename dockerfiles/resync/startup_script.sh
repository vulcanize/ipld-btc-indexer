#!/bin/sh
# Runs the db migrations and starts the super node services

# Exit if the variable tests fail
set -e
set +x

# Check the database variables are set
test $DATABASE_HOSTNAME
test $DATABASE_NAME
test $DATABASE_PORT
test $DATABASE_USER
test $DATABASE_PASSWORD
test $IPFS_INIT
set +e

# Export our database variables so that the IPFS Postgres plugin can use them
export IPFS_PGHOST=$DATABASE_HOSTNAME
export IPFS_PGUSER=$DATABASE_USER
export IPFS_PGDATABASE=$DATABASE_NAME
export IPFS_PGPORT=$DATABASE_PORT
export IPFS_PGPASSWORD=$DATABASE_PASSWORD

# If IPFS_INIT is true
if [[ "$IPFS_INIT" = true ]] ; then
    # initialize PG-IPFS
    echo "Initializing Postgres-IPFS profile"
    ./ipfs init --profile=postgresds
else
    echo "IPFS profile already initialized, skipping initialization"
fi

# If IPFS initialization was successful
if [[ $? -eq 0 ]]; then
    echo "Beginning the vulcanizedb super node process"
    ./vulcanizedb resync --config=config.toml 2>&1 | tee -a vulcanizedb.log &
else
    echo "Could not initialize IPFS."
    exit 1
fi

# If Vulcanizedb startup was successful
if [ $? -eq 0 ]; then
    echo "Resync successfully booted"
else
    echo "Could not start vulcanizedb resync process. Is the config file correct?"
    exit 1
fi

tail -f vulcanizedb.log