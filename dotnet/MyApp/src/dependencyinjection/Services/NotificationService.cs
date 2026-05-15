namespace DependencyInjection.Services;

public class NotificationService : INotificationService
{
    private readonly IMessageService _messageService;
    private readonly IUserService _userService;

    public NotificationService(IMessageService messageService, IUserService userService)
    {
        _messageService = messageService;
        _userService = userService;
    }

    public void NotifyUser(string message)
    {
        var user = _userService.GetCurrentUser();
        var formattedMessage = $"[{user}] {message}";
        _messageService.SendMessage(formattedMessage);
    }
}