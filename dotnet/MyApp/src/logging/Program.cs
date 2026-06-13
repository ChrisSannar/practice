using System;
using System.Collections.Generic;
using Microsoft.Extensions.Logging;

Console.WriteLine("=== Logging & Error Handling Demo ===\n");

var loggerFactory = LoggerFactory.Create(builder =>
{
    builder.AddConsole();
    builder.SetMinimumLevel(LogLevel.Information);
});

ILogger<Program> logger = loggerFactory.CreateLogger<Program>();

logger.LogInformation("Application starting up");
logger.LogDebug("Debug information - detailed flow");
logger.LogWarning("Warning: configuration file not found, using defaults");
logger.LogError("Error: failed to connect to database");
logger.LogCritical("Critical: system running out of memory");

Console.WriteLine("\n--- Logging Levels Demo ---");
LogAtEachLevel(loggerFactory.CreateLogger("Levels"));

Console.WriteLine("\n--- Error Handling Demo ---");
var calculator = new Calculator();
try
{
    var result = calculator.Divide(10, 0);
}
catch (DivideByZeroException ex)
{
    logger.LogError(ex, "Caught divide by zero");
    Console.WriteLine($"Caught: {ex.Message}");
}

Console.WriteLine("\n--- Global Exception Handler Demo ---");
var errorHandler = new ErrorHandler(logger);
errorHandler.HandleError("Test error 1");
errorHandler.HandleError("Test error 2");

Console.WriteLine("\n--- Custom Logger Demo ---");
var customLogger = new CustomLogger();
customLogger.Log("INFO", "This is an info message");
customLogger.Log("WARNING", "This is a warning message");
customLogger.Log("ERROR", "This is an error message");

Console.WriteLine("\n--- Try-Catch-Finally Demo ---");
var resourceManager = new ResourceManager();
try
{
    resourceManager.Open();
    resourceManager.Process();
}
catch (Exception ex)
{
    Console.WriteLine($"Exception: {ex.Message}");
}
finally
{
    resourceManager.Close();
    Console.WriteLine("Cleanup completed in finally block");
}

Console.WriteLine("\n=== Demo Complete ===");

Console.WriteLine("\n=== Running Tests ===\n");

int passed = 0;
int failed = 0;

void RunTest(string name, Action test)
{
    try
    {
        test();
        Console.WriteLine($"[PASS] {name}");
        passed++;
    }
    catch (Exception ex)
    {
        Console.WriteLine($"[FAIL] {name}: {ex.Message}");
        failed++;
    }
}

RunTest("Logger_CreatesLoggerOfType", () =>
{
    var factory = LoggerFactory.Create(_ => {});
    var logger = factory.CreateLogger<Program>();
    Assert.NotNull(logger);
});

RunTest("Logger_LogsAtInformationLevel", () =>
{
    var messages = new List<string>();
    var factory = LoggerFactory.Create(builder =>
    {
        builder.AddProvider(new TestLoggerProvider(m => messages.Add(m)));
    });
    var logger = factory.CreateLogger<Program>();
    logger.LogInformation("Test message");
    Assert.True(messages.Any(m => m.Contains("Test message")));
});

RunTest("Logger_LogWarning_Works", () =>
{
    var messages = new List<string>();
    var factory = LoggerFactory.Create(builder =>
    {
        builder.AddProvider(new TestLoggerProvider(m => messages.Add(m)));
    });
    var logger = factory.CreateLogger<Program>();
    logger.LogWarning("Warning message");
    Assert.True(messages.Any(m => m.Contains("Warning message")));
});

RunTest("ErrorHandler_StoresErrors", () =>
{
    var testLogger = new TestLogger();
    var handler = new ErrorHandler(testLogger);
    handler.HandleError("Error 1");
    handler.HandleError("Error 2");
    Assert.Equal(2, testLogger.ErrorCount);
});

RunTest("Calculator_DivideByZero_Throws", () =>
{
    var calc = new Calculator();
    try
    {
        calc.Divide(10, 0);
        throw new Exception("Should have thrown");
    }
    catch (DivideByZeroException)
    {
    }
});

RunTest("ResourceManager_CleanupInFinally", () =>
{
    var manager = new ResourceManager();
    try
    {
        manager.Open();
        manager.Process();
    }
    catch { }
    finally
    {
        manager.Close();
    }
    Assert.True(manager.IsClosed);
});

Console.WriteLine($"\n=== Test Results: {passed} passed, {failed} failed ===");

static void LogAtEachLevel(ILogger logger)
{
    logger.LogTrace("Trace - most detailed");
    logger.LogDebug("Debug - detailed info");
    logger.LogInformation("Information - general info");
    logger.LogWarning("Warning - something unusual");
    logger.LogError("Error - something failed");
    logger.LogCritical("Critical - system in trouble");
}

class Calculator
{
    public int Divide(int a, int b)
    {
        if (b == 0) throw new DivideByZeroException("Cannot divide by zero");
        return a / b;
    }
}

class ErrorHandler
{
    private readonly ILogger _logger;
    public ErrorHandler(ILogger logger) => _logger = logger;
    public void HandleError(string message) => _logger.LogError(message);
}

class CustomLogger
{
    public void Log(string level, string message)
    {
        Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] [{level}] {message}");
    }
}

class ResourceManager
{
    public bool IsClosed { get; private set; }
    public void Open() => Console.WriteLine("Opening resource");
    public void Process() => Console.WriteLine("Processing");
    public void Close() => IsClosed = true;
}

class TestLoggerProvider : ILoggerProvider
{
    private readonly Action<string> _onLog;
    public TestLoggerProvider(Action<string> onLog) => _onLog = onLog;
    public ILogger CreateLogger(string categoryName) => new TestLogger(_onLog, categoryName);
    public void Dispose() { }
}

class TestLogger : ILogger
{
    private readonly Action<string> _onLog;
    public string CategoryName { get; }
    public int ErrorCount { get; private set; }

    public TestLogger() { }
    public TestLogger(Action<string> onLog, string categoryName)
    {
        _onLog = onLog;
        CategoryName = categoryName;
    }

    public IDisposable? BeginScope<TState>(TState state) where TState : notnull => null;

    public bool IsEnabled(LogLevel logLevel) => true;

    public void Log<TState>(LogLevel logLevel, EventId eventId, TState state, Exception? exception, Func<TState, Exception?, string> formatter)
    {
        var message = formatter(state, exception);
        _onLog?.Invoke(message);
        if (logLevel >= LogLevel.Error) ErrorCount++;
    }
}

static class Assert
{
    public static void True(bool condition) 
    {
        if (!condition) throw new Exception("Expected true");
    }
    public static void NotNull(object? obj)
    {
        if (obj == null) throw new Exception("Expected non-null");
    }
    public static void Equal<T>(T expected, T actual)
    {
        if (!EqualityComparer<T>.Default.Equals(expected, actual))
            throw new Exception($"Expected {expected}, got {actual}");
    }
}