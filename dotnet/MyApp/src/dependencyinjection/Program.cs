using Microsoft.Extensions.DependencyInjection;
using DependencyInjection.Services;
using Xunit;

Console.WriteLine("=== Dependency Injection Demo ===\n");

var services = new ServiceCollection();

services.AddSingleton<ISingletonService, SingletonService>();
services.AddScoped<IScopedService, ScopedService>();
services.AddTransient<ITransientService, TransientService>();

services.AddSingleton<IMessageService>(new MessageService("CustomPrefix"));
services.AddScoped<IUserService, UserService>();
services.AddTransient<INotificationService, NotificationService>();

var serviceProvider = services.BuildServiceProvider();

Console.WriteLine("--- Service Lifetimes Demo ---");
using var scope1 = serviceProvider.CreateScope();
var s1 = scope1.ServiceProvider.GetRequiredService<ISingletonService>();
var s2 = scope1.ServiceProvider.GetRequiredService<ISingletonService>();
Console.WriteLine($"Singleton 1: {s1.GetInfo()}");
Console.WriteLine($"Singleton 2: {s2.GetInfo()}");
Console.WriteLine($"Same instance? {s1.Id == s2.Id}");

var sc1 = scope1.ServiceProvider.GetRequiredService<IScopedService>();
var sc2 = scope1.ServiceProvider.GetRequiredService<IScopedService>();
Console.WriteLine($"Scoped 1: {sc1.GetInfo()}");
Console.WriteLine($"Scoped 2: {sc2.GetInfo()}");
Console.WriteLine($"Same instance? {sc1.Id == sc2.Id}");

var t1 = scope1.ServiceProvider.GetRequiredService<ITransientService>();
var t2 = scope1.ServiceProvider.GetRequiredService<ITransientService>();
Console.WriteLine($"Transient 1: {t1.GetInfo()}");
Console.WriteLine($"Transient 2: {t2.GetInfo()}");
Console.WriteLine($"Same instance? {t1.Id == t2.Id}");

Console.WriteLine("\n--- New Scope ---");
using var scope2 = serviceProvider.CreateScope();
var s3 = scope2.ServiceProvider.GetRequiredService<ISingletonService>();
var sc3 = scope2.ServiceProvider.GetRequiredService<IScopedService>();
Console.WriteLine($"Singleton across scopes: {s3.Id} (same as before: {s3.Id == s1.Id})");
Console.WriteLine($"Scoped new scope: {sc3.Id} (different from {sc1.Id}: {sc3.Id != sc1.Id})");

Console.WriteLine("\n--- Chained Dependencies Demo ---");
var notifier = serviceProvider.GetRequiredService<INotificationService>();
notifier.NotifyUser("System maintenance at 5PM");

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
        Console.WriteLine($"[FAIL] {name}: {ex.InnerException?.Message ?? ex.Message}");
        failed++;
    }
}

RunTest("AddSingleton_SameInstanceReturned", () =>
{
    var svcs = new ServiceCollection();
    svcs.AddSingleton<ISingletonService, SingletonService>();
    var provider = svcs.BuildServiceProvider();
    var i1 = provider.GetRequiredService<ISingletonService>();
    var i2 = provider.GetRequiredService<ISingletonService>();
    Assert.Same(i1, i2);
});

RunTest("AddTransient_NewInstanceEachTime", () =>
{
    var svcs = new ServiceCollection();
    svcs.AddTransient<ITransientService, TransientService>();
    var provider = svcs.BuildServiceProvider();
    var i1 = provider.GetRequiredService<ITransientService>();
    var i2 = provider.GetRequiredService<ITransientService>();
    Assert.NotSame(i1, i2);
});

RunTest("AddScoped_SameInstanceWithinScope", () =>
{
    var svcs = new ServiceCollection();
    svcs.AddScoped<IScopedService, ScopedService>();
    var provider = svcs.BuildServiceProvider();
    using var scope = provider.CreateScope();
    var i1 = scope.ServiceProvider.GetRequiredService<IScopedService>();
    var i2 = scope.ServiceProvider.GetRequiredService<IScopedService>();
    Assert.Same(i1, i2);
});

RunTest("AddScoped_DifferentInstanceAcrossScopes", () =>
{
    var svcs = new ServiceCollection();
    svcs.AddScoped<IScopedService, ScopedService>();
    var provider = svcs.BuildServiceProvider();
    using var scope1 = provider.CreateScope();
    using var scope2 = provider.CreateScope();
    var i1 = scope1.ServiceProvider.GetRequiredService<IScopedService>();
    var i2 = scope2.ServiceProvider.GetRequiredService<IScopedService>();
    Assert.NotSame(i1, i2);
});

RunTest("ConstructorInjection_ResolvesChainedDependencies", () =>
{
    var svcs = new ServiceCollection();
    svcs.AddSingleton<IMessageService, MessageService>();
    svcs.AddScoped<IUserService, UserService>();
    svcs.AddTransient<INotificationService, NotificationService>();
    var provider = svcs.BuildServiceProvider();
    var ns = provider.GetRequiredService<INotificationService>();
    Assert.NotNull(ns);
});

RunTest("AddSingleton_WithInstance_UsesProvidedInstance", () =>
{
    var custom = new MessageService("TestPrefix");
    var svcs = new ServiceCollection();
    svcs.AddSingleton<IMessageService>(custom);
    var provider = svcs.BuildServiceProvider();
    var resolved = provider.GetRequiredService<IMessageService>();
    Assert.Same(custom, resolved);
    Assert.Equal("TestPrefix: Hello from MessageService!", resolved.GetMessage());
});

RunTest("GetService_ReturnsNullWhenNotRegistered", () =>
{
    var svcs = new ServiceCollection();
    var provider = svcs.BuildServiceProvider();
    var result = provider.GetService<IMessageService>();
    Assert.Null(result);
});

RunTest("GetRequiredService_ThrowsWhenNotRegistered", () =>
{
    var svcs = new ServiceCollection();
    var provider = svcs.BuildServiceProvider();
    Assert.Throws<InvalidOperationException>(() => provider.GetRequiredService<IMessageService>());
});

Console.WriteLine($"\n=== Test Results: {passed} passed, {failed} failed ===");