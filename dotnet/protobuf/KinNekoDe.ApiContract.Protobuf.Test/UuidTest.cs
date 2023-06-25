using Kinnekode.Protobuf;
using Uuid = Kinnekode.Protobuf.Uuid;

namespace KinNekoDe.ApiContract.Protobuf.Test;

public class UuidTest
{
    public class NotNullableTestData : TheoryData<Guid, string>
    {
        public NotNullableTestData()
        {
            Add(new Guid("ED420000-0000-0000-0000-000000000000"), "ed420000-0000-0000-0000-000000000000");
            Add(new Guid("ed420000-0000-0000-0000-000000000000"), "ed420000-0000-0000-0000-000000000000");
        }
    }

    public class InvalidTestData : TheoryData<string>
    {
        public InvalidTestData()
        {
            Add(string.Empty);
            Add("12");
            Add("ED420000000000000000000000000000");
            Add("ed420000000000000000000000000000");
            Add("NoGuid");
        }
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void FromGuid(Guid input, string expected)
    {
        var actual = Uuid.FromGuid(input);

        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }


    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ToGuid(Guid expected, string input)
    {
        var message = new Uuid() { Value = input };

        var actual = message.ToGuid();

        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void ToGuid_Invalid(string input)
    {
        var message = new Uuid() { Value = input };

        Assert.Throws<FormatException>(() => message.ToGuid());
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void TryParseToDecimal(Guid expected, string input)
    {
        var message = new Uuid() { Value = input };

        var parseable = message.TryParseToGuid(out var actual);

        Assert.True(parseable);
        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void TryParseToDecimal_Invalid(string input)
    {
        var message = new Uuid() { Value = input };

        var parseable = message.TryParseToGuid(out var actual);

        Assert.False(parseable);
        Assert.Equal(Guid.Empty, actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ExplicitOperator_Guid(Guid input, string expected)
    {
        var actual = (Uuid)input;
        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }

    [Fact]
    public void ExplicitOperator_NullableGuid()
    {
        var input = (Guid?)null;
        var actual = (Uuid?)input;
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void ToNullableGuid(Guid expected, string input)
    {
        var message = new Uuid() { Value = input };

        Guid? actual = message.ToNullableGuid();
        Assert.Equal(expected, actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void ToNullableGuid_Invalid(string input)
    {
        var message = new Uuid() { Value = input };

        Assert.Throws<FormatException>(() => message.ToNullableGuid());
    }

    [Fact]
    public void ToNullableGuid_MessageNull()
    {
        Uuid? message = null;

        Guid? actual = message.ToNullableGuid();
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void TryParseToNullableGuid(Guid expected, string input)
    {
        var message = new Uuid() { Value = input };

        var parseable = message.TryParseToNullableGuid(out var actual);

        Assert.True(parseable);
        Assert.Equal(expected, actual);
    }
        
    [Fact]
    public void TryParseToNullableGuid_Null()
    {
        Uuid? message = null;

        var parseable = message.TryParseToNullableGuid(out var actual);
        Assert.True(parseable);
        Assert.Null(actual);
    }

    [Theory]
    [ClassData(typeof(InvalidTestData))]
    public void TryParseToNullableGuid_Invalid(string input)
    {
        var message = new Uuid() { Value = input };

        var parseable = message.TryParseToNullableGuid(out var actual);

        Assert.False(parseable);
        Assert.Equal(Guid.Empty, actual);
    }

    [Theory]
    [ClassData(typeof(NotNullableTestData))]
    public void FromNullableGuid(Guid? input, string expected)
    {
        var actual = UuidExtensions.FromNullableGuid(input);
        Assert.NotNull(actual);
        Assert.Equal(expected, actual.Value);
    }

    [Fact]
    public void FromNullableGuid_Null()
    {
        Guid? input = null;

        var actual = UuidExtensions.FromNullableGuid(input);
        Assert.Null(actual);
    }
}