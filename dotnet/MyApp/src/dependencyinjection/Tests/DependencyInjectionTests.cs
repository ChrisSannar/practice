using Microsoft.Extensions.DependencyInjection;
using DependencyInjection.Services;
using Xunit;

namespace DependencyInjection.Tests;

public class DependencyInjectionTests
{
    [Fact]
    public void AddSingleton_SameInstanceReturned()
    {
        var services = new ServiceCollection();
        services.AddSingleton<ISingletonService, SingletonService>();

        var provider = services.BuildServiceProvider();

        var instance1 = provider.GetRequiredService<ISingletonService>();
        var instance2 = provider.GetRequiredService<ISingletonService>();

        Assert.Same(instance1, instance2);
    }

    [Fact]
    public void AddTransient_NewInstanceEachTime()
    {
        var services = new ServiceCollection();
        services.AddTransient<ITransientService, TransientService>();

        var provider = services.BuildServiceProvider();

        var instance1 = provider.GetRequiredService<ITransientService>();
        var instance2 = provider.GetRequiredService<ITransientService>();

        Assert.NotSame(instance1, instance2);
    }

    [Fact]
    public void AddScoped_SameInstanceWithinScope()
    {
        var services = new ServiceCollection();
        services.AddScoped<IScopedService, ScopedService>();

        var provider = services.BuildServiceProvider();

        using var scope = provider.CreateScope();
        var instance1 = scope.ServiceProvider.GetRequiredService<IScopedService>();
        var instance2 = scope.ServiceProvider.GetRequiredService<IScopedService>();

        Assert.Same(instance1, instance2);
    }

    [Fact]
    public void AddScoped_DifferentInstanceAcrossScopes()
    {
        var services = new ServiceCollection();
        services.AddScoped<IScopedService, ScopedService>();

        var provider = services.BuildServiceProvider();

        using var scope1 = provider.CreateScope();
        using var scope2 = provider.CreateScope();
        var instance1 = scope1.ServiceProvider.GetRequiredService<IScopedService>();
        var instance2 = scope2.ServiceProvider.GetRequiredService<IScopedService>();

        Assert.NotSame(instance1, instance2);
    }

    [Fact]
    public void ConstructorInjection_ResolvesChainedDependencies()
    {
        var services = new ServiceCollection();
        services.AddSingleton<IMessageService, MessageService>();
        services.AddScoped<IUserService, UserService>();
        services.AddTransient<INotificationService, NotificationService>();

        var provider = services.BuildServiceProvider();

        var notificationService = provider.GetRequiredService<INotificationService>();

        Assert.NotNull(notificationService);
    }

    [Fact]
    public void AddSingleton_WithInstance_UsesProvidedInstance()
    {
        var customMessageService = new MessageService("TestPrefix");
        var services = new ServiceCollection();
        services.AddSingleton<IMessageService>(customMessageService);

        var provider = services.BuildServiceProvider();

        var resolved = provider.GetRequiredService<IMessageService>();

        Assert.Same(customMessageService, resolved);
        Assert.Equal("TestPrefix: Hello from MessageService!", resolved.GetMessage());
    }

    [Fact]
    public void GetService_ReturnsNullWhenNotRegistered()
    {
        var services = new ServiceCollection();
        var provider = services.BuildServiceProvider();

        var result = provider.GetService<IMessageService>();

        Assert.Null(result);
    }

    [Fact]
    public void GetRequiredService_ThrowsWhenNotRegistered()
    {
        var services = new ServiceCollection();
        var provider = services.BuildServiceProvider();

        Assert.Throws<InvalidOperationException>(() =>
            provider.GetRequiredService<IMessageService>());
    }

    [Fact]
    public void FactoryRegistration_CreatesInstancesCorrectly()
    {
        var services = new ServiceCollection();
        services.AddTransient<ITransientService>(_ => new TransientService());

        var provider = services.BuildServiceProvider();

        var instance = provider.GetRequiredService<ITransientService>();

        Assert.NotNull(instance);
    }

    [Fact]
    public void MultipleServices_SameLifetime_ShareInstances()
    {
        var services = new ServiceCollection();
        services.AddSingleton<IMessageService, MessageService>();
        services.AddSingleton<IUserService, UserService>();

        var provider = services.BuildServiceProvider();

        var msgService = provider.GetRequiredService<IMessageService>();
        var userService = provider.GetRequiredService<IUserService>();

        Assert.NotNull(msgService);
        Assert.NotNull(userService);
    }
}