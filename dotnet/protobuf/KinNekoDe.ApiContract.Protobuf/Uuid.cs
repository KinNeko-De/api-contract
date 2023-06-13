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
        return Guid.Parse(Value);
    }

    public bool TryParseToGuid(out Guid guid)
    {
        if (Guid.TryParse(Value, out Guid parsedGuid))
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