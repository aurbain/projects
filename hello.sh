#!/bin/bash
# This script will say hello
echo "hello"

# --- Prime Numbers up to 10000 using Sieve of Eratosthenes ---
LIMIT=10000
# Create a boolean array (or map equivalent) representing numbers up to LIMIT
# true means potentially prime, false means composite
isPrime=()
for (( i=0; i<=LIMIT; i++ )); do
    isPrime[$i]=true
done

isPrime[0]=false
isPrime[1]=false

p=2
while [ $p -le $LIMIT ]; do
    # If isPrime[p] is true, then it is a prime
    if [ "$isPrime[$p]" = "true" ]; then
        echo $p # Print the prime number
        # Mark all multiples of p starting from p*p as not prime
        for (( i=p*p; i<=LIMIT; i+=p )); do
            isPrime[$i]=false
        done
    fi
    ((p++))
done
