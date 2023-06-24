using Google.Protobuf;

namespace Kinnekode.Protobuf;


public partial class Decimal : ICustomDiagnosticMessage
{
    public static Decimal FromDecimal(decimal value)
    {
        return new Decimal { Value = ToMessageValue(value) };
    }

    public decimal ToDecimal()
    {
        return decimal.Parse(Value);
    }

    public bool TryParseToDecimal(out decimal value)
    {
        if (decimal.TryParse(Value, out decimal parsedDecimal))
        {
            value = parsedDecimal;
            return true;
        }

        value = decimal.Zero;
        return false;
    }

    public static explicit operator Decimal(decimal value) => FromDecimal(value);
    public static explicit operator Decimal?(decimal? value) => DecimalExtensions.FromNullableDecimal(value);

    string ICustomDiagnosticMessage.ToDiagnosticString()
    {
        return Value;
    }

    private static string ToMessageValue(decimal value)
    {
        return value.ToString("0.00", System.Globalization.CultureInfo.InvariantCulture);
    }
}