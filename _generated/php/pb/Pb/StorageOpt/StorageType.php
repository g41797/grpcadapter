<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: station.proto

namespace Pb\StorageOpt;

use UnexpectedValueException;

/**
 * Protobuf type <code>pb.StorageOpt.StorageType</code>
 */
class StorageType
{
    /**
     * Generated from protobuf enum <code>Disk = 0;</code>
     */
    const Disk = 0;
    /**
     * Generated from protobuf enum <code>Memory = 1;</code>
     */
    const Memory = 1;

    private static $valueToName = [
        self::Disk => 'Disk',
        self::Memory => 'Memory',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(StorageType::class, \Pb\StorageOpt_StorageType::class);

