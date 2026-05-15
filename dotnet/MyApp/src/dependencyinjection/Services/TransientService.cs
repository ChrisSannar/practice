namespace DependencyInjection.Services;

public class TransientService : ITransientService
{
    public Guid Id { get; } = Guid.NewGuid();

    public string GetInfo() => $"Transient (new instance): {Id}";
}