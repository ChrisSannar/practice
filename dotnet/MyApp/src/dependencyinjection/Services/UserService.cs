namespace DependencyInjection.Services;

public class UserService : IUserService
{
    private readonly IMessageService _messageService;

    public UserService(IMessageService messageService)
    {
        _messageService = messageService;
    }

    public string GetCurrentUser() => "Admin";

    public int GetUserCount() => 42;
}