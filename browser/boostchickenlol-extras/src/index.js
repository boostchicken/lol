var conf = []
function cache() {
    fetch("https://lol.boostchicken.dev/liveconfig")
    .then((response) => response.json())
    .then((data) => conf = data)
    setInterval(cache,60)
}   
chrome.omnibox.onInputChanged.addListener((text, suggest) => {
    if(conf.length == 0) {
         cache()
    }
    let results =[]
    conf.Entries.forEach(entry => {
        let command = entry.command
        let description = entry.description
        let value = entry.value
        if(command.contains(text) || description.contains(text)) {
            results.push({
                    content: `<mark>${text}</mark>`,
                    deletable: false,
                    description: `<mark>${command}</mark> - ${description}, url: <url>${value}</url> `
                })
        }
    });
    suggest(results)
});
chrome.omnibox.onInputEntered.addListener((text, OnInputEnteredDisposition) => {
        if (OnInputEnteredDisposition === 'currentTab') {
            chrome.tabs.create({text});
        } else {
            chrome.tabs.update({text});
        }
    });
chrome.omnibox.setDefaultSuggestion({
    description: 'Enter text to search command names and metadata'
});


