namespace DependencyInjection.Services;

public interface IScopedService
{
    Guid Id { get; }
    string GetInfo();
}