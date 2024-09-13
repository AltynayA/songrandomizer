

let parentcontainer = document.getElementById("parent-container")

async function getPlaylist() {
    const url = "http://localhost:3000/songs"
    try {
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error(`Response status: ${response.status}`);
        }
    
        let json = await response.json();
        console.log(json)
        return json
      } catch (error) {
        console.error(error.message);
        return "[]"
      }
}


async function getRandom() {
    var songs = await getPlaylist()
    console.log(songs)
    x = songs.length
    console.log(x)
    console.log(songs[0])
    let ind = Math.floor(Math.random() * x)
    let resultsong = songs[ind]
    //console.log(resultsong)
    songname = document.getElementById("song")
    songname.innerText = resultsong.title
    author = document.getElementById("author")
    author.innerText = resultsong.author
    link = document.getElementById("songlink")
    link.href = resultsong.spotify_link
    //console.log(link.href)
    
}
//let button = document.getElementById("random-button")

//button.addEventListener("click", getRandom())


