namespace DependencyInjection.Services;

public interface IUserService
{
    string GetCurrentUser();
    int GetUserCount();
}