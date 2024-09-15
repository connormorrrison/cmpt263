"""
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
"""


# Calculate coin combinations
def calculate_combinations(amount, coins, current_combinations, current_index):
    # Base case
    if amount == 0:
        display_combinations(current_combinations)
        return None
    
    # Else we have more than 0 for the amount, use recursion
    for i in range(current_index, len(coins)):
        if coins[i] <= amount:
            # Subtracts coins[i] quantity from amount
            # Appends coins[i] to current_combinations
            calculate_combinations(amount - coins[i], coins, current_combinations + [coins[i]], i)


# Display coin combinations
def display_combinations(current_combinations):
    result = {}
    for coin in current_combinations:
        result[coin] = result.get(coin, 0) + 1
    
    # Print in format
    formatted_strings = []
    for coin, count in result.items():
        formatted_string = f"{count}@{coin}"
        formatted_strings.append(formatted_string)
    
    output = ", ".join(formatted_strings)
    print(output)


def main():
    amount = int(input(""))
    coins = [200, 100, 50, 25, 10, 5, 1]

    # Special case where amount = 0
    if amount == 0:
        print("0@0")
        return None

    # Calculate results
    calculate_combinations(amount, coins, current_combinations=[], current_index=0)


if __name__ == "__main__":
    main()
