﻿using System.Globalization;
using Google.Protobuf;

namespace Kinnekode.Protobuf;
public partial class Decimal : ICustomDiagnosticMessage
{
    public static Decimal FromDecimal(decimal value)
    {
        return new Decimal() { Value = ToMessageValue(value) };
    }

    public decimal ToDecimal()
    {
        return decimal.Parse(Value, GetNumberStyles(), CultureInfo.InvariantCulture);
    }

    public bool TryParseToDecimal(out decimal value)
    {
        var parseable = decimal.TryParse(Value, GetNumberStyles(), CultureInfo.InvariantCulture, out decimal parsedDecimal);
        value = parseable ? parsedDecimal : default;
        return parseable;
    }

    string ICustomDiagnosticMessage.ToDiagnosticString()
    {
        return Value;
    }

    private static string ToMessageValue(decimal value)
    {
        return value.ToString(CultureInfo.InvariantCulture);
    }

    private static NumberStyles GetNumberStyles()
    {
        return NumberStyles.AllowLeadingSign | NumberStyles.AllowDecimalPoint;
    }

    public static explicit operator Decimal(decimal value) => FromDecimal(value);
    public static explicit operator Decimal?(decimal? value) => DecimalExtensions.FromNullableDecimal(value);
}