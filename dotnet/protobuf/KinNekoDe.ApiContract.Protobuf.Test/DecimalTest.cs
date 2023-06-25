using Kinnekode.Protobuf;
using Decimal = Kinnekode.Protobuf.Decimal;

namespace KinNekoDe.ApiContract.Protobuf.Test;

public class DecimalTest
{
    public class NotNullableTestData : TheoryData<decimal, string>
    {
        public NotNullableTestData()
        {
            Add(5.124m, "5.124");
            Add(-5.124m, "-5.124");
            Add(5m, "5");
            Add(-5m, "-5");
            Add(0m, "0");
        }
    }

    public class InvalidTestData : TheoryData<string>
    {
        public InvalidTestData()
        {
            Add("                  5.124");
            Add("5.124                  ");
            Add("5.124+");
            Add("5.124-");
            Add("5,124");
            Add("NoNumber");
        }
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void FromDecimal(decimal input, string expected)
    {
        var actual = Decimal.FromDecimal(input);

        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }


    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ToDecimal(decimal expected, string input)
    {
        var message = new Decimal() { Value = input };

        var actual = message.ToDecimal();

        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void ToDecimal_Invalid(string input)
    {
        var message = new Decimal() { Value = input };

        Assert.Throws<FormatException>(() => message.ToDecimal());
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void TryParseToDecimal(decimal expected, string input)
    {
        var message = new Decimal() { Value = input };

        var parseable = message.TryParseToDecimal(out var actual);

        Assert.True(parseable);
        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void TryParseToDecimal_Invalid(string input)
    {
        var message = new Decimal() { Value = input };

        var parseable = message.TryParseToDecimal(out var actual);

        Assert.False(parseable);
        Assert.Equal(decimal.Zero, actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ExplicitOperator_Decimal(decimal input, string expected)
    {
        var actual = (Decimal)input;
        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }

    [Fact]
    public void ExplicitOperator_NullableDecimal()
    {
        var input = (decimal?)null;
        var actual = (Decimal?)input;
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ToNullableDecimal(decimal expected, string input)
    {
        var message = new Decimal() { Value = input };

        decimal? actual = message.ToNullableDecimal();
        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void ToNullableDecimal_Invalid(string input)
    {
        var message = new Decimal() { Value = input };

        Assert.Throws<FormatException>(() => message.ToNullableDecimal());
    }

    [Fact]
    public void ToNullableDecimal_MessageNull()
    {
        Decimal? message = null;

        decimal? actual = message.ToNullableDecimal();
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void TryParseToNullableDecimal(decimal expected, string input)
    {
        var message = new Decimal() { Value = input };

        var parseable = message.TryParseToNullableDecimal(out var actual);

        Assert.True(parseable);
        Assert.Equal(expected, actual);
    }
        
    [Fact]
    public void TryParseToNullableDecimal_Null()
    {
        Decimal? message = null;

        var parseable = message.TryParseToNullableDecimal(out var actual);
        Assert.True(parseable);
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void TryParseToNullableDecimal_Invalid(string input)
    {
        var message = new Decimal() { Value = input };

        var parseable = message.TryParseToNullableDecimal(out var actual);

        Assert.False(parseable);
        Assert.Equal(decimal.Zero, actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void FromNullableDecimal(decimal? input, string expected)
    {
        var actual = DecimalExtensions.FromNullableDecimal(input);
        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }

    [Fact]
    public void FromNullableDecimal_Null()
    {
        decimal? input = null;

        var actual = DecimalExtensions.FromNullableDecimal(input);
        Assert.Null(actual);
    }
}