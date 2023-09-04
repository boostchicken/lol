
// Mock the window object
const mockWindow = {
  onload: jest.fn(),
};

// Mock the Object.defineProperties method
Object.defineProperties = jest.fn();

describe('index.ts', () => {
  beforeEach(() => {
    // Reset the mock functions before each test
    jest.clearAllMocks();
  });

  it('should define properties for scheme and host', () => {
    // Set up the test environment
    global.window = mockWindow;

    // Call the onload function
    window.onload();

    // Assert that the Object.defineProperties method was called with the correct arguments
    expect(Object.defineProperties).toHaveBeenCalledWith(conf, {
      scheme: {
        set: expect.any(Function),
      },
      host: {
        set: expect.any(Function),
      },
    });
  });
});