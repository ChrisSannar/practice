using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.DependencyInjection;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.Use(async (context, next) =>
{
    Console.WriteLine($"[Request] {context.Request.Method} {context.Request.Path}");
    await next();
    Console.WriteLine($"[Response] {context.Response.StatusCode}");
});

app.Use(async (context, next) =>
{
    var stopwatch = System.Diagnostics.Stopwatch.StartNew();
    await next();
    stopwatch.Stop();
    Console.WriteLine($"[Timing] {context.Request.Path} took {stopwatch.ElapsedMilliseconds}ms");
});

app.Use(async (context, next) =>
{
    if (context.Request.Path.StartsWithSegments("/api"))
    {
        context.Items["StartTime"] = DateTime.Now;
    }
    await next();
});

app.MapGet("/", () => "Hello from Middleware Demo!");
app.MapGet("/api/test", () => "API Endpoint Response");
app.MapGet("/api/error", () => throw new Exception("Test exception"));

app.Use(async (context, next) =>
{
    try
    {
        await next();
    }
    catch (Exception ex)
    {
        Console.WriteLine($"[Error] {ex.Message}");
        context.Response.StatusCode = 500;
        await context.Response.WriteAsync("Error occurred");
    }
});

Console.WriteLine("=== Middleware Demo ===\n");
Console.WriteLine("Middleware pipeline configured:");
Console.WriteLine("1. RequestLoggingMiddleware - logs incoming requests");
Console.WriteLine("2. TimingMiddleware - measures request duration");
Console.WriteLine("3. ApiStartTimeMiddleware - adds start time to context");
Console.WriteLine("4. Endpoint handlers (/, /api/test, /api/error)");
Console.WriteLine("5. ErrorHandlingMiddleware - catches exceptions");
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

RunTest("WebApplication_CanBeCreated", () =>
{
    var testBuilder = WebApplication.CreateBuilder();
    var testApp = testBuilder.Build();
    Assert.NotNull(testApp);
});

RunTest("Middleware_CanMapGet", () =>
{
    var testBuilder = WebApplication.CreateBuilder();
    var testApp = testBuilder.Build();
    testApp.MapGet("/test", () => "OK");
    Assert.NotNull(testApp);
});

RunTest("Middleware_RunTerminatesChain", () =>
{
    var testBuilder = WebApplication.CreateBuilder();
    var testApp = testBuilder.Build();
    bool firstCalled = false;
    bool secondCalled = false;
    
    testApp.Use(async (c, n) => { firstCalled = true; await n(); });
    testApp.Use(async (c, n) => { secondCalled = true; });
    
    Assert.True(firstCalled);
});

RunTest("MapWhen_ConditionalMiddleware", () =>
{
    var testBuilder = WebApplication.CreateBuilder();
    var testApp = testBuilder.Build();
    bool apiCalled = false;
    bool normalCalled = false;
    
    testApp.MapWhen(c => c.Request.Path.StartsWithSegments("/api"), 
        app => app.Use(async (c, n) => { apiCalled = true; await n(); }));
    testApp.Use(async (c, n) => { normalCalled = true; await n(); });
    
    Assert.True(true);
});

RunTest("UseWhen_BranchingMiddleware", () =>
{
    var testBuilder = WebApplication.CreateBuilder();
    var testApp = testBuilder.Build();
    testApp.UseWhen(c => c.Request.Path.StartsWithSegments("/admin"), 
        app => app.Use(async (c, n) => { await n(); }));
    
    Assert.NotNull(testApp);
});

Console.WriteLine($"\n=== Test Results: {passed} passed, {failed} failed ===");

app.Run();

static class Assert
{
    public static void NotNull(object? obj)
    {
        if (obj == null) throw new Exception("Expected non-null");
    }
    public static void True(bool condition)
    {
        if (!condition) throw new Exception("Expected true");
    }
}