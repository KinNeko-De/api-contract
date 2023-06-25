using System;
using Google.Protobuf;
using Google.Protobuf.WellKnownTypes;

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
        var parseable = Guid.TryParseExact(Value, "D", out Guid parsedGuid);
        guid = parseable ? parsedGuid : default;
        return parseable;
    }

    string ICustomDiagnosticMessage.ToDiagnosticString()
    {
        return Value;
    }

    public static explicit operator Uuid(Guid guid) => FromGuid(guid);
    public static explicit operator Uuid?(Guid? guid) => UuidExtensions.FromNullableGuid(guid);
}