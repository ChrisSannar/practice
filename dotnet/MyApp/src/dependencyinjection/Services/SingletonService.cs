namespace DependencyInjection.Services;

public class SingletonService : ISingletonService
{
    public Guid Id { get; } = Guid.NewGuid();

    public string GetInfo() => $"Singleton (same instance): {Id}";
}