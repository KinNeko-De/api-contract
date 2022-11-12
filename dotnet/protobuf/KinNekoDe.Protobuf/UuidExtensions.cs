using System;

namespace Kinnekode.Protobuf;

public static class UuidExtensions
{
    public static Guid? ToNullableGuid(this Uuid? uuid)
    {
        return uuid?.ToGuid();
    }

    public static bool TryParseToNullableGuid(this Uuid? uuid, out Guid? nullableGuid)
    {
        if (uuid == null)
        {
            nullableGuid = null;
            return true;
        }

        var result = uuid.TryParseToGuid(out Guid guid);
        nullableGuid = guid;
        return result;
    }

    public static Uuid? FromNullableGuid(Guid? guid)
    {
        return guid == null ? null : new Uuid() { Value = guid.Value.ToString("D") };
    }
}