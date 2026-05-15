namespace DependencyInjection.Services;

public interface ISingletonService
{
    Guid Id { get; }
    string GetInfo();
}