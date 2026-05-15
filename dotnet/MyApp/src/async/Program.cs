using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Async.Examples;

Console.WriteLine("=== Async/Await Demo ===\n");

var dataService = new AsyncDataService();
var parallelAsync = new ParallelAsync();
var configureAwaitExample = new ConfigureAwaitExample();

Console.WriteLine("--- Basic Async Methods ---");
var data = await dataService.GetDataAsync();
Console.WriteLine($"GetDataAsync result: {data}");

var count = await dataService.GetCountAsync();
Console.WriteLine($"GetCountAsync result: {count}");

var items = await dataService.GetItemsAsync();
Console.WriteLine($"GetItemsAsync result: {string.Join(", ", items)}");

Console.WriteLine("\n--- Task.WhenAll (Parallel Execution) ---");
var allResults = await parallelAsync.FetchAllAsync();
Console.WriteLine($"FetchAllAsync result: {allResults}");

Console.WriteLine("\n--- Task.WhenAny (First Completed) ---");
var firstResult = await parallelAsync.FetchFirstCompletedAsync();
Console.WriteLine($"FetchFirstCompletedAsync result: {firstResult}");

Console.WriteLine("\n--- ConfigureAwait(false) ---");
await configureAwaitExample.ProcessAsync();
Console.WriteLine("ProcessAsync completed with ConfigureAwait(false)");

Console.WriteLine("\n--- CancellationToken Support ---");
using var cts = new CancellationTokenSource();
cts.CancelAfter(50);
try
{
    await configureAwaitExample.ProcessWithCancellationAsync(cts.Token);
}
catch (OperationCanceledException)
{
    Console.WriteLine("CancellationToken: Operation was cancelled as expected");
}

Console.WriteLine("\n=== Demo Complete ===");

Console.WriteLine("\n=== Running Tests ===\n");

int passed = 0;
int failed = 0;

void RunTest(string name, Func<Task> test)
{
    try
    {
        test().GetAwaiter().GetResult();
        Console.WriteLine($"[PASS] {name}");
        passed++;
    }
    catch (Exception ex)
    {
        Console.WriteLine($"[FAIL] {name}: {ex.InnerException?.Message ?? ex.Message}");
        failed++;
    }
}

RunTest("AsyncMethod_ReturnsTask", async () =>
{
    var result = await dataService.GetDataAsync();
    Assert.Equal("Data loaded", result);
});

RunTest("TaskWhenAll_WaitsForAll", async () =>
{
    var result = await parallelAsync.FetchAllAsync();
    Assert.Contains("User1", result);
    Assert.Contains("User2", result);
    Assert.Contains("User3", result);
});

RunTest("TaskWhenAny_ReturnsFirst", async () =>
{
    var result = await parallelAsync.FetchFirstCompletedAsync();
    Assert.StartsWith("User", result);
});

RunTest("CancellationToken_ThrowsOnCancel", async () =>
{
    using var cts = new CancellationTokenSource();
    cts.CancelAfter(10);
    try
    {
        await configureAwaitExample.ProcessWithCancellationAsync(cts.Token);
        throw new Exception("Should have thrown");
    }
    catch (OperationCanceledException)
    {
    }
});

RunTest("CancellationToken_ReturnsOnComplete", async () =>
{
    using var cts = new CancellationTokenSource();
    cts.CancelAfter(1000);
    var result = await configureAwaitExample.CalculateWithTokenAsync(cts.Token);
    Assert.Equal(42, result);
});

Console.WriteLine($"\n=== Test Results: {passed} passed, {failed} failed ===");

static class Assert
{
    public static void Equal<T>(T expected, T actual) 
    {
        if (!EqualityComparer<T>.Default.Equals(expected, actual))
            throw new Exception($"Expected {expected}, got {actual}");
    }
    public static void Contains(string substring, string value)
    {
        if (!value.Contains(substring))
            throw new Exception($"Expected '{value}' to contain '{substring}'");
    }
    public static void StartsWith(string prefix, string value)
    {
        if (!value.StartsWith(prefix))
            throw new Exception($"Expected '{value}' to start with '{prefix}'");
    }
}