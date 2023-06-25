using System;
using Google.Protobuf;

namespace Kinnekode.Protobuf;

public partial class Uuid : ICustomDiagnosticMessage
{
    public static Uuid FromGuid(Guid guid)
    {
        return new Uuid() { Value = guid.ToString("D") };
    }

    public Guid ToGuid()
    {
        return Guid.ParseExact(Value, "D");
    }

    public bool TryParseToGuid(out Guid guid)
    {
        if (Guid.TryParseExact(Value, "D", out Guid parsedGuid))
        {
            guid = parsedGuid;
            return true;
        }

        guid = Guid.Empty;
        return false;
    }

    public static explicit operator Uuid(Guid guid) => FromGuid(guid);
    public static explicit operator Uuid?(Guid? guid) => UuidExtensions.FromNullableGuid(guid);

    string ICustomDiagnosticMessage.ToDiagnosticString()
    {
        return Value;
    }
}