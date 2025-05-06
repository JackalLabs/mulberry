// SPDX-License-Identifier: MIT
pragma solidity <0.9.0 >=0.6.2 ^0.8.0 ^0.8.20 ^0.8.26;

// lib/foundry-chainlink-toolkit/src/interfaces/feeds/AggregatorV3Interface.sol

interface AggregatorV3Interface {
  function decimals() external view returns (uint8);
  function description() external view returns (string memory);
  function version() external view returns (uint256);
  function getRoundData(uint80 _roundId) external view returns (
    uint80 roundId,
    int256 answer,
    uint256 startedAt,
    uint256 updatedAt,
    uint80 answeredInRound
  );
  function latestRoundData() external view returns (
    uint80 roundId,
    int256 answer,
    uint256 startedAt,
    uint256 updatedAt,
    uint80 answeredInRound
  );
}

// lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/utils/Context.sol

// OpenZeppelin Contracts (last updated v5.0.1) (utils/Context.sol)

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    function _contextSuffixLength() internal view virtual returns (uint256) {
        return 0;
    }
}

// lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/utils/math/Math.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/math/Math.sol)

/**
 * @dev Standard math utilities missing in the Solidity language.
 */
library Math {
    /**
     * @dev Muldiv operation overflow.
     */
    error MathOverflowedMulDiv();

    enum Rounding {
        Floor, // Toward negative infinity
        Ceil, // Toward positive infinity
        Trunc, // Toward zero
        Expand // Away from zero
    }

    /**
     * @dev Returns the addition of two unsigned integers, with an overflow flag.
     */
    function tryAdd(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            uint256 c = a + b;
            if (c < a) return (false, 0);
            return (true, c);
        }
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, with an overflow flag.
     */
    function trySub(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b > a) return (false, 0);
            return (true, a - b);
        }
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, with an overflow flag.
     */
    function tryMul(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
            // benefit is lost if 'b' is also tested.
            // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
            if (a == 0) return (true, 0);
            uint256 c = a * b;
            if (c / a != b) return (false, 0);
            return (true, c);
        }
    }

    /**
     * @dev Returns the division of two unsigned integers, with a division by zero flag.
     */
    function tryDiv(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b == 0) return (false, 0);
            return (true, a / b);
        }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers, with a division by zero flag.
     */
    function tryMod(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b == 0) return (false, 0);
            return (true, a % b);
        }
    }

    /**
     * @dev Returns the largest of two numbers.
     */
    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? a : b;
    }

    /**
     * @dev Returns the smallest of two numbers.
     */
    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two numbers. The result is rounded towards
     * zero.
     */
    function average(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b) / 2 can overflow.
        return (a & b) + (a ^ b) / 2;
    }

    /**
     * @dev Returns the ceiling of the division of two numbers.
     *
     * This differs from standard division with `/` in that it rounds towards infinity instead
     * of rounding towards zero.
     */
    function ceilDiv(uint256 a, uint256 b) internal pure returns (uint256) {
        if (b == 0) {
            // Guarantee the same behavior as in a regular Solidity division.
            return a / b;
        }

        // (a + b - 1) / b can overflow on addition, so we distribute.
        return a == 0 ? 0 : (a - 1) / b + 1;
    }

    /**
     * @notice Calculates floor(x * y / denominator) with full precision. Throws if result overflows a uint256 or
     * denominator == 0.
     * @dev Original credit to Remco Bloemen under MIT license (https://xn--2-umb.com/21/muldiv) with further edits by
     * Uniswap Labs also under MIT license.
     */
    function mulDiv(uint256 x, uint256 y, uint256 denominator) internal pure returns (uint256 result) {
        unchecked {
            // 512-bit multiply [prod1 prod0] = x * y. Compute the product mod 2^256 and mod 2^256 - 1, then use
            // use the Chinese Remainder Theorem to reconstruct the 512 bit result. The result is stored in two 256
            // variables such that product = prod1 * 2^256 + prod0.
            uint256 prod0 = x * y; // Least significant 256 bits of the product
            uint256 prod1; // Most significant 256 bits of the product
            assembly {
                let mm := mulmod(x, y, not(0))
                prod1 := sub(sub(mm, prod0), lt(mm, prod0))
            }

            // Handle non-overflow cases, 256 by 256 division.
            if (prod1 == 0) {
                // Solidity will revert if denominator == 0, unlike the div opcode on its own.
                // The surrounding unchecked block does not change this fact.
                // See https://docs.soliditylang.org/en/latest/control-structures.html#checked-or-unchecked-arithmetic.
                return prod0 / denominator;
            }

            // Make sure the result is less than 2^256. Also prevents denominator == 0.
            if (denominator <= prod1) {
                revert MathOverflowedMulDiv();
            }

            ///////////////////////////////////////////////
            // 512 by 256 division.
            ///////////////////////////////////////////////

            // Make division exact by subtracting the remainder from [prod1 prod0].
            uint256 remainder;
            assembly {
                // Compute remainder using mulmod.
                remainder := mulmod(x, y, denominator)

                // Subtract 256 bit number from 512 bit number.
                prod1 := sub(prod1, gt(remainder, prod0))
                prod0 := sub(prod0, remainder)
            }

            // Factor powers of two out of denominator and compute largest power of two divisor of denominator.
            // Always >= 1. See https://cs.stackexchange.com/q/138556/92363.

            uint256 twos = denominator & (0 - denominator);
            assembly {
                // Divide denominator by twos.
                denominator := div(denominator, twos)

                // Divide [prod1 prod0] by twos.
                prod0 := div(prod0, twos)

                // Flip twos such that it is 2^256 / twos. If twos is zero, then it becomes one.
                twos := add(div(sub(0, twos), twos), 1)
            }

            // Shift in bits from prod1 into prod0.
            prod0 |= prod1 * twos;

            // Invert denominator mod 2^256. Now that denominator is an odd number, it has an inverse modulo 2^256 such
            // that denominator * inv = 1 mod 2^256. Compute the inverse by starting with a seed that is correct for
            // four bits. That is, denominator * inv = 1 mod 2^4.
            uint256 inverse = (3 * denominator) ^ 2;

            // Use the Newton-Raphson iteration to improve the precision. Thanks to Hensel's lifting lemma, this also
            // works in modular arithmetic, doubling the correct bits in each step.
            inverse *= 2 - denominator * inverse; // inverse mod 2^8
            inverse *= 2 - denominator * inverse; // inverse mod 2^16
            inverse *= 2 - denominator * inverse; // inverse mod 2^32
            inverse *= 2 - denominator * inverse; // inverse mod 2^64
            inverse *= 2 - denominator * inverse; // inverse mod 2^128
            inverse *= 2 - denominator * inverse; // inverse mod 2^256

            // Because the division is now exact we can divide by multiplying with the modular inverse of denominator.
            // This will give us the correct result modulo 2^256. Since the preconditions guarantee that the outcome is
            // less than 2^256, this is the final result. We don't need to compute the high bits of the result and prod1
            // is no longer required.
            result = prod0 * inverse;
            return result;
        }
    }

    /**
     * @notice Calculates x * y / denominator with full precision, following the selected rounding direction.
     */
    function mulDiv(uint256 x, uint256 y, uint256 denominator, Rounding rounding) internal pure returns (uint256) {
        uint256 result = mulDiv(x, y, denominator);
        if (unsignedRoundsUp(rounding) && mulmod(x, y, denominator) > 0) {
            result += 1;
        }
        return result;
    }

    /**
     * @dev Returns the square root of a number. If the number is not a perfect square, the value is rounded
     * towards zero.
     *
     * Inspired by Henry S. Warren, Jr.'s "Hacker's Delight" (Chapter 11).
     */
    function sqrt(uint256 a) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        // For our first guess, we get the biggest power of 2 which is smaller than the square root of the target.
        //
        // We know that the "msb" (most significant bit) of our target number `a` is a power of 2 such that we have
        // `msb(a) <= a < 2*msb(a)`. This value can be written `msb(a)=2**k` with `k=log2(a)`.
        //
        // This can be rewritten `2**log2(a) <= a < 2**(log2(a) + 1)`
        // → `sqrt(2**k) <= sqrt(a) < sqrt(2**(k+1))`
        // → `2**(k/2) <= sqrt(a) < 2**((k+1)/2) <= 2**(k/2 + 1)`
        //
        // Consequently, `2**(log2(a) / 2)` is a good first approximation of `sqrt(a)` with at least 1 correct bit.
        uint256 result = 1 << (log2(a) >> 1);

        // At this point `result` is an estimation with one bit of precision. We know the true value is a uint128,
        // since it is the square root of a uint256. Newton's method converges quadratically (precision doubles at
        // every iteration). We thus need at most 7 iteration to turn our partial result with one bit of precision
        // into the expected uint128 result.
        unchecked {
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            return min(result, a / result);
        }
    }

    /**
     * @notice Calculates sqrt(a), following the selected rounding direction.
     */
    function sqrt(uint256 a, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = sqrt(a);
            return result + (unsignedRoundsUp(rounding) && result * result < a ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 2 of a positive value rounded towards zero.
     * Returns 0 if given 0.
     */
    function log2(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 128;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 64;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 32;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 16;
            }
            if (value >> 8 > 0) {
                value >>= 8;
                result += 8;
            }
            if (value >> 4 > 0) {
                value >>= 4;
                result += 4;
            }
            if (value >> 2 > 0) {
                value >>= 2;
                result += 2;
            }
            if (value >> 1 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 2, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log2(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log2(value);
            return result + (unsignedRoundsUp(rounding) && 1 << result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 10 of a positive value rounded towards zero.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10 ** 64) {
                value /= 10 ** 64;
                result += 64;
            }
            if (value >= 10 ** 32) {
                value /= 10 ** 32;
                result += 32;
            }
            if (value >= 10 ** 16) {
                value /= 10 ** 16;
                result += 16;
            }
            if (value >= 10 ** 8) {
                value /= 10 ** 8;
                result += 8;
            }
            if (value >= 10 ** 4) {
                value /= 10 ** 4;
                result += 4;
            }
            if (value >= 10 ** 2) {
                value /= 10 ** 2;
                result += 2;
            }
            if (value >= 10 ** 1) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 10, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log10(value);
            return result + (unsignedRoundsUp(rounding) && 10 ** result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 256 of a positive value rounded towards zero.
     * Returns 0 if given 0.
     *
     * Adding one to the result gives the number of pairs of hex symbols needed to represent `value` as a hex string.
     */
    function log256(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 16;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 8;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 4;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 2;
            }
            if (value >> 8 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 256, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log256(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log256(value);
            return result + (unsignedRoundsUp(rounding) && 1 << (result << 3) < value ? 1 : 0);
        }
    }

    /**
     * @dev Returns whether a provided rounding mode is considered rounding up for unsigned integers.
     */
    function unsignedRoundsUp(Rounding rounding) internal pure returns (bool) {
        return uint8(rounding) % 2 == 1;
    }
}

// lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/utils/math/SignedMath.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/math/SignedMath.sol)

/**
 * @dev Standard signed math utilities missing in the Solidity language.
 */
library SignedMath {
    /**
     * @dev Returns the largest of two signed numbers.
     */
    function max(int256 a, int256 b) internal pure returns (int256) {
        return a > b ? a : b;
    }

    /**
     * @dev Returns the smallest of two signed numbers.
     */
    function min(int256 a, int256 b) internal pure returns (int256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two signed numbers without overflow.
     * The result is rounded towards zero.
     */
    function average(int256 a, int256 b) internal pure returns (int256) {
        // Formula from the book "Hacker's Delight"
        int256 x = (a & b) + ((a ^ b) >> 1);
        return x + (int256(uint256(x) >> 255) & (a ^ b));
    }

    /**
     * @dev Returns the absolute unsigned value of a signed value.
     */
    function abs(int256 n) internal pure returns (uint256) {
        unchecked {
            // must be unchecked in order to support `n = type(int256).min`
            return uint256(n >= 0 ? n : -n);
        }
    }
}

// lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/access/Ownable.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/Ownable.sol)

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * The initial owner is set to the address provided by the deployer. This can
 * later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    /**
     * @dev The caller account is not authorized to perform an operation.
     */
    error OwnableUnauthorizedAccount(address account);

    /**
     * @dev The owner is not a valid owner account. (eg. `address(0)`)
     */
    error OwnableInvalidOwner(address owner);

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the address provided by the deployer as the initial owner.
     */
    constructor(address initialOwner) {
        if (initialOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(initialOwner);
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        if (owner() != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        if (newOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}

// lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/utils/Strings.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/Strings.sol)

/**
 * @dev String operations.
 */
library Strings {
    bytes16 private constant HEX_DIGITS = "0123456789abcdef";
    uint8 private constant ADDRESS_LENGTH = 20;

    /**
     * @dev The `value` string doesn't fit in the specified `length`.
     */
    error StringsInsufficientHexLength(uint256 value, uint256 length);

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = Math.log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            /// @solidity memory-safe-assembly
            assembly {
                ptr := add(buffer, add(32, length))
            }
            while (true) {
                ptr--;
                /// @solidity memory-safe-assembly
                assembly {
                    mstore8(ptr, byte(mod(value, 10), HEX_DIGITS))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    /**
     * @dev Converts a `int256` to its ASCII `string` decimal representation.
     */
    function toStringSigned(int256 value) internal pure returns (string memory) {
        return string.concat(value < 0 ? "-" : "", toString(SignedMath.abs(value)));
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation.
     */
    function toHexString(uint256 value) internal pure returns (string memory) {
        unchecked {
            return toHexString(value, Math.log256(value) + 1);
        }
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation with fixed length.
     */
    function toHexString(uint256 value, uint256 length) internal pure returns (string memory) {
        uint256 localValue = value;
        bytes memory buffer = new bytes(2 * length + 2);
        buffer[0] = "0";
        buffer[1] = "x";
        for (uint256 i = 2 * length + 1; i > 1; --i) {
            buffer[i] = HEX_DIGITS[localValue & 0xf];
            localValue >>= 4;
        }
        if (localValue != 0) {
            revert StringsInsufficientHexLength(value, length);
        }
        return string(buffer);
    }

    /**
     * @dev Converts an `address` with fixed length of 20 bytes to its not checksummed ASCII `string` hexadecimal
     * representation.
     */
    function toHexString(address addr) internal pure returns (string memory) {
        return toHexString(uint256(uint160(addr)), ADDRESS_LENGTH);
    }

    /**
     * @dev Returns true if the two strings are equal.
     */
    function equal(string memory a, string memory b) internal pure returns (bool) {
        return bytes(a).length == bytes(b).length && keccak256(bytes(a)) == keccak256(bytes(b));
    }
}

// src/Jackal.sol

abstract contract Jackal {
    event PostedFile(address from, string merkle, uint64 size, string note, uint64 expires);
    event BoughtStorage(address from, string for_address, uint64 duration_days, uint64 size_bytes, string referral);
    event DeletedFile(address from, string merkle, uint64 start);
    event RequestedReportForm(address from, string prover, string merkle, string owner, uint64 start);
    event PostedKey(address from, string key);
    event DeletedFileTree(address from, string hash_path, string account);
    event ProvisionedFileTree(address from, string editors, string viewers, string tracking_number);
    event PostedFileTree(
        address from,
        string account,
        string hash_parent,
        string hash_child,
        string contents,
        string viewers,
        string editors,
        string tracking_number
    );
    event AddedViewers(address from, string viewer_ids, string viewer_keys, string for_address, string file_owner);
    event RemovedViewers(address from, string viewer_ids, string for_address, string file_owner);
    event ResetViewers(address from, string for_address, string file_owner);
    event ChangedOwner(address from, string for_address, string file_owner, string new_owner);
    event AddedEditors(address from, string editor_ids, string editor_keys, string for_address, string file_owner);
    event RemovedEditors(address from, string editor_ids, string for_address, string file_owner);
    event ResetEditors(address from, string for_address, string file_owner);
    event CreatedNotification(address from, string to, string contents, string private_contents);
    event DeletedNotification(address from, string notification_from, uint64 time);
    event BlockedSenders(address from, string[] to_block);

    struct JackalMessage {
        string id;
        address sender;
        uint256 height;
        uint256 value;
    }

    JackalMessage[] public messages;

    function newMessage(address sender, string memory messageType) private returns (string memory) {
        uint256 height = block.number;

        // concat messageType and address
        string memory s = string.concat(messageType, Strings.toHexString(uint160(sender), 20));
        s = string.concat(s, Strings.toString(height));

        JackalMessage memory message = JackalMessage(s, sender, height, msg.value);

        messages.push(message);

        return s;
    }

    function _remove(uint256 index) internal {
        require(index < messages.length);
        messages[index] = messages[messages.length - 1];
        messages.pop();
    }

    function refund(string memory id) public {
        refundFrom(payable(msg.sender), id);
    }

    function refundFrom(address payable from, string memory id) public hasAllowance(from) {
        for (uint256 i = 0; i < messages.length; i++) {
            JackalMessage memory m = messages[i];
            if (Strings.equal(m.id, id)) {
                // if we found the item we're looking for
                if (m.height + 60 < block.number) {
                    // if 60 blocks have passed and this message is still here
                    from.transfer(m.value);
                    _remove(i);
                }
            }
        }
    }

    function getPrice() public view virtual returns (int256);

    mapping(address => mapping(address => bool)) public allowances;

    modifier validAddress() {
        require(msg.sender != address(0), "Invalid sender address");
        _;
    }

    modifier hasAllowance(address from) {
        require(getAllowance(msg.sender, from), "No allowance set for contract");
        _;
    }

    function getAllowance(address from, address to) public view returns (bool) {
        if (from == to) {
            return true;
        }
        return allowances[to][from];
    }

    function addAllowance(address allow) public {
        allowances[msg.sender][allow] = true;
    }

    function removeAllowance(address allow) public {
        allowances[msg.sender][allow] = false;
    }

    function getStoragePrice(uint64 filesize, uint256 months) public view returns (uint256) {
        // requires chainlink oracle, will not work on localnet
        uint256 price = uint256(getPrice());
        uint256 storagePrice = 15; // price at 8 decimal places
        uint256 multiplier = 2;

        uint256 fs = filesize;
        if (fs <= 1024 * 1024) {
            fs = 1024 * 1024; // minimum file size of one MB for pricing
        }

        // Calculate the price in wei
        // 1e8 adjusts for the 8 decimals of USD, 1e18 converts ETH to wei
        uint256 BSM = storagePrice * multiplier * months * fs;
        uint256 p = (BSM * 1e8 * 1e18) / (price * 1099511627776);

        if (p == 0) {
            p = 5000 gwei;
        }

        return p;
    }

    function postFile(string memory merkle, uint64 filesize, string memory note, uint64 expires) public payable {
        postFileFrom(msg.sender, merkle, filesize, note, expires);
    }

    function postFileFrom(address from, string memory merkle, uint64 filesize, string memory note, uint64 expires)
        public
        payable
        validAddress
        hasAllowance(from)
    {
        require(expires >= 30 || expires == 0);
        if (expires != 0) {
            uint256 pE = getStoragePrice(filesize, 2400); // 12 * 200 months
            require(msg.value >= pE, string.concat("Insufficient payment, need ", Strings.toString(pE), " wei"));
        }
        newMessage(from, "PostedFile");
        emit PostedFile(from, merkle, filesize, note, expires);
    }

    function buyStorage(string memory for_address, uint64 duration_days, uint64 size_bytes, string memory referral)
        public
        payable
    {
        buyStorageFrom(msg.sender, for_address, duration_days, size_bytes, referral);
    }

    function buyStorageFrom(
        address from,
        string memory for_address,
        uint64 duration_days,
        uint64 size_bytes,
        string memory referral
    ) public payable validAddress hasAllowance(from) {
        uint256 pE = getStoragePrice(size_bytes, duration_days / 30); // months
        require(msg.value >= pE, string.concat("Insufficient payment, need ", Strings.toString(pE), " wei"));
        newMessage(from, "BoughtStorage");
        emit BoughtStorage(from, for_address, duration_days, size_bytes, referral);
    }

    function deleteFile(string memory merkle, uint64 start) public {
        deleteFileFrom(msg.sender, merkle, start);
    }

    function deleteFileFrom(address from, string memory merkle, uint64 start) public validAddress hasAllowance(from) {
        newMessage(from, "DeletedFile");
        emit DeletedFile(from, merkle, start); // file deletion is free
    }

    function requestReportForm(string memory prover, string memory merkle, string memory owner, uint64 start) public {
        requestReportFormFrom(msg.sender, prover, merkle, owner, start);
    }

    function requestReportFormFrom(
        address from,
        string memory prover,
        string memory merkle,
        string memory owner,
        uint64 start
    ) public validAddress hasAllowance(from) {
        newMessage(from, "RequestedReportForm");
        emit RequestedReportForm(from, prover, merkle, owner, start);
    }

    function postKey(string memory key) public {
        postKeyFrom(msg.sender, key);
    }

    function postKeyFrom(address from, string memory key) public validAddress hasAllowance(from) {
        newMessage(from, "PostedKey");
        emit PostedKey(from, key);
    }

    function deleteFileTree(string memory hash_path, string memory account) public {
        deleteFileTreeFrom(msg.sender, hash_path, account);
    }

    function deleteFileTreeFrom(address from, string memory hash_path, string memory account)
        public
        validAddress
        hasAllowance(from)
    {
        newMessage(from, "DeletedFileTree");
        emit DeletedFileTree(from, hash_path, account);
    }

    function provisionFileTree(string memory editors, string memory viewers, string memory tracking_number) public {
        provisionFileTreeFrom(msg.sender, editors, viewers, tracking_number);
    }

    function provisionFileTreeFrom(
        address from,
        string memory editors,
        string memory viewers,
        string memory tracking_number
    ) public validAddress hasAllowance(from) {
        newMessage(from, "ProvisionedFileTree");
        emit ProvisionedFileTree(from, editors, viewers, tracking_number);
    }

    function postFileTree(
        string memory account,
        string memory hash_parent,
        string memory hash_child,
        string memory contents,
        string memory viewers,
        string memory editors,
        string memory tracking_number
    ) public {
        postFileTreeFrom(msg.sender, account, hash_parent, hash_child, contents, viewers, editors, tracking_number);
    }

    function postFileTreeFrom(
        address from,
        string memory account,
        string memory hash_parent,
        string memory hash_child,
        string memory contents,
        string memory viewers,
        string memory editors,
        string memory tracking_number
    ) public validAddress hasAllowance(from) {
        newMessage(from, "PostedFileTree");
        emit PostedFileTree(from, account, hash_parent, hash_child, contents, viewers, editors, tracking_number);
    }

    function addViewers(
        string memory viewer_ids,
        string memory viewer_keys,
        string memory for_address,
        string memory file_owner
    ) public {
        addViewersFrom(msg.sender, viewer_ids, viewer_keys, for_address, file_owner);
    }

    function addViewersFrom(
        address from,
        string memory viewer_ids,
        string memory viewer_keys,
        string memory for_address,
        string memory file_owner
    ) public validAddress hasAllowance(from) {
        newMessage(from, "AddedViewers");
        emit AddedViewers(from, viewer_ids, viewer_keys, for_address, file_owner);
    }

    function removeViewers(string memory viewer_ids, string memory for_address, string memory file_owner) public {
        removeViewersFrom(msg.sender, viewer_ids, for_address, file_owner);
    }

    function removeViewersFrom(
        address from,
        string memory viewer_ids,
        string memory for_address,
        string memory file_owner
    ) public validAddress hasAllowance(from) {
        newMessage(from, "RemovedViewers");
        emit RemovedViewers(from, viewer_ids, for_address, file_owner);
    }

    function resetViewers(string memory for_address, string memory file_owner) public {
        resetViewersFrom(msg.sender, for_address, file_owner);
    }

    function resetViewersFrom(address from, string memory for_address, string memory file_owner)
        public
        validAddress
        hasAllowance(from)
    {
        newMessage(from, "ResetViewers");
        emit ResetViewers(from, for_address, file_owner);
    }

    function changeOwner(string memory for_address, string memory file_owner, string memory new_owner) public {
        changeOwnerFrom(msg.sender, for_address, file_owner, new_owner);
    }

    function changeOwnerFrom(address from, string memory for_address, string memory file_owner, string memory new_owner)
        public
        validAddress
        hasAllowance(from)
    {
        newMessage(from, "ChangedOwner");
        emit ChangedOwner(from, for_address, file_owner, new_owner);
    }

    function addEditors(
        string memory editor_ids,
        string memory editor_keys,
        string memory for_address,
        string memory file_owner
    ) public {
        addEditorsFrom(msg.sender, editor_ids, editor_keys, for_address, file_owner);
    }

    function addEditorsFrom(
        address from,
        string memory editor_ids,
        string memory editor_keys,
        string memory for_address,
        string memory file_owner
    ) public validAddress hasAllowance(from) {
        newMessage(from, "AddedEditors");
        emit AddedEditors(from, editor_ids, editor_keys, for_address, file_owner);
    }

    function removeEditors(string memory editor_ids, string memory for_address, string memory file_owner) public {
        removeEditorsFrom(msg.sender, editor_ids, for_address, file_owner);
    }

    function removeEditorsFrom(
        address from,
        string memory editor_ids,
        string memory for_address,
        string memory file_owner
    ) public validAddress hasAllowance(from) {
        newMessage(from, "RemovedEditors");
        emit RemovedEditors(from, editor_ids, for_address, file_owner);
    }

    function resetEditors(string memory for_address, string memory file_owner) public {
        resetEditorsFrom(msg.sender, for_address, file_owner);
    }

    function resetEditorsFrom(address from, string memory for_address, string memory file_owner)
        public
        validAddress
        hasAllowance(from)
    {
        newMessage(from, "ResetEditors");
        emit ResetEditors(from, for_address, file_owner);
    }

    function createNotification(string memory to, string memory contents, string memory private_contents) public {
        createNotificationFrom(msg.sender, to, contents, private_contents);
    }

    function createNotificationFrom(
        address from,
        string memory to,
        string memory contents,
        string memory private_contents
    ) public validAddress hasAllowance(from) {
        newMessage(from, "CreatedNotification");
        emit CreatedNotification(from, to, contents, private_contents);
    }

    function deleteNotification(string memory notification_from, uint64 time) public {
        deleteNotificationFrom(msg.sender, notification_from, time);
    }

    function deleteNotificationFrom(address from, string memory notification_from, uint64 time)
        public
        validAddress
        hasAllowance(from)
    {
        newMessage(from, "DeletedNotification");
        emit DeletedNotification(from, notification_from, time);
    }

    function blockSenders(string[] memory to_block) public {
        blockSendersFrom(msg.sender, to_block);
    }

    function blockSendersFrom(address from, string[] memory to_block) public validAddress hasAllowance(from) {
        newMessage(from, "BlockedSenders");
        emit BlockedSenders(from, to_block);
    }
}

// src/JackalV1.sol

contract JackalBridge is Ownable, Jackal {
    AggregatorV3Interface internal priceFeed;

    address[] public relays;

    constructor(address[] memory _relays, address _priceFeed) Ownable(msg.sender) {
        require(_relays.length > 0, "must provide relays");

        priceFeed = AggregatorV3Interface(_priceFeed);
        relays = _relays;
    }

    // Modifier to restrict access to owner or relays
    modifier onlyOwnerOrRelay() {
        require(msg.sender == owner() || isRelay(msg.sender), "not owner or relay");
        _;
    }

    function finishMessage(string memory id) public onlyOwnerOrRelay { // needs to be from a relayer
        for (uint i = 0; i < messages.length; i ++) {
            JackalMessage memory m = messages[i];
            if (Strings.equal(m.id, id)) { // if we found the item we're looking for
                _remove(i);
            }
        }
    }

    function isRelay(address _relay) internal view returns (bool) {
        for (uint256 i = 0; i < relays.length; i++) {
            if (relays[i] == _relay) {
                return true;
            }
        }
        return false;
    }

    // Function to add a relay, only callable by the owner
    function addRelay(address _relay) public onlyOwner {
        relays.push(_relay);
    }

    // Function to remove a relay, only callable by the owner
    function removeRelay(address _relay) public onlyOwner {
        require(relays.length > 1); // require there to be at least one relay in the list after removal

        for (uint256 i = 0; i < relays.length; i++) {
            if (relays[i] == _relay) {
                relays[i] = relays[relays.length - 1];
                relays.pop();
                break;
            }
        }
    }

    function distributeBalance() public onlyOwnerOrRelay {
        uint256 balance = address(this).balance;
        require(balance > 2, "not enough wei to distribute");

        uint256 ownerShare = balance / 2;
        payable(owner()).transfer(ownerShare);

        uint256 relayShare = balance - ownerShare; // Remaining 50%
        uint256 perRelay = relayShare / relays.length;

        for (uint256 i = 0; i < relays.length; i++) {
            payable(relays[i]).transfer(perRelay);
        }
    }

    function getPrice() public view override returns (int256) {
        (, int256 price,,,) = priceFeed.latestRoundData();
        return price;
    }
}

