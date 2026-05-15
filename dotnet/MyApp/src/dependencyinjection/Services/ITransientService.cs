namespace DependencyInjection.Services;

public interface ITransientService
{
    Guid Id { get; }
    string GetInfo();
}