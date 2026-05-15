namespace Async.Examples;

public class ParallelAsync
{
    public async Task<string> FetchAllAsync()
    {
        var tasks = new List<Task<string>>
        {
            GetUserAsync(1),
            GetUserAsync(2),
            GetUserAsync(3)
        };

        var results = await Task.WhenAll(tasks);
        return string.Join(", ", results);
    }

    public async Task<string> FetchFirstCompletedAsync()
    {
        var tasks = new List<Task<string>>
        {
            GetUserAsync(1),
            GetUserAsync(2),
            GetUserAsync(3)
        };

        var completedTask = await Task.WhenAny(tasks);
        return await completedTask;
    }

    private async Task<string> GetUserAsync(int id)
    {
        await Task.Delay(100);
        return $"User{id}";
    }
}