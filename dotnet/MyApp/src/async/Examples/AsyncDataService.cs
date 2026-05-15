namespace Async.Examples;

public class AsyncDataService
{
    public async Task<string> GetDataAsync()
    {
        await Task.Delay(100);
        return "Data loaded";
    }

    public async Task<int> GetCountAsync()
    {
        await Task.Delay(50);
        return 42;
    }

    public async Task<List<string>> GetItemsAsync()
    {
        await Task.Delay(75);
        return new List<string> { "Item1", "Item2", "Item3" };
    }
}