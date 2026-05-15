namespace DependencyInjection.Services;

public class ScopedService : IScopedService
{
    public Guid Id { get; } = Guid.NewGuid();

    public string GetInfo() => $"Scoped (per request): {Id}";
}