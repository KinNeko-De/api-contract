namespace Kinnekode.Protobuf;

public static class DecimalExtensions
{
    public static decimal? ToNullableDecimal(this Decimal? message)
    {
        return message?.ToDecimal();
    }

    public static bool TryParseToNullableDecimal(this Decimal? message, out decimal? nullableDecimal)
    {
        if (message == null)
        {
            nullableDecimal = null;
            return true;
        }

        var result = message.TryParseToDecimal(out decimal value);
        nullableDecimal = value;
        return result;
    }

    public static Decimal? FromNullableDecimal(decimal? value)
    {
        return value == null ? null : Decimal.FromDecimal(value.Value);
    }
}