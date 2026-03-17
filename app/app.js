const rawInput = document.getElementById('input');
const inputButton = document.getElementById('btn-input');
const tategaki = document.querySelector('.tategaki');
const kakikudashi = document.querySelector('.kakikudashi');

async function renderTategaki(sentence) {
    const container = document.getElementById("tategaki");
    container.innerHTML = ""; 

    characters = sentence.Characters
    characters.forEach((char, index) => {
        const box = document.createElement("div");
        box.className = "kanji-box";
        box.innerHTML = `
<span>${char.Kanji}</span>
<button class="btn-issaidoku" onclick="updateSaidoku(${index})">再</button>
<div class="line-s" id="saidoku${index}"></div>
<button class="btn-isjuku" onclick="updateJuku(${index})">熟</button>
<div class="line-j" id="juku${index}"></div>
<textarea class="okurigana2" onchange="updateSentence(this.value, ${index}, 'okuri2')"></textarea>
<textarea class="kaeriten" onchange="updateSentence(this.value, ${index}, 'kaeri')"></textarea>
<textarea class="okurigana" onchange="updateSentence(this.value, ${index}, 'okuri')"></textarea>
`;
        box.dataset.index = index; 
        container.appendChild(box);
    });}

inputButton.addEventListener('click', async () => {
    const text = rawInput.value;
    if (!text) return; 
    kakikudashi.innerHTML = text;

    try {
        const response = await fetch('/api/sentences', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ text: text })
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const sentenceData = await response.json();
        renderTategaki(sentenceData.sentence);

    } catch (error) {
        console.error('Error:', error);
        alert("Failure; check the console");
    }
});

async function updateSaidoku(index) {
    const line = document.getElementById('saidoku' + index);
    if (line.style.display == 'none') {
        line.style.display = 'block';
    } else {
        line.style.display = 'none';
    }
    updateSentence('', index, 'saidoku');
}
async function updateJuku(index) {
    const line = document.getElementById('juku' + index);
    if (line.style.display == 'none') {
        line.style.display = 'block';
    } else {
        line.style.display = 'none';
    }
    updateSentence('', index, 'juku');
}

async function updateSentence(value, index, type) {
    try {
        const response = await fetch('/api/sentences/1', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ index: index, text: value, sentence_id: 1 , type: type})
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const kakikudashibun = await response.json();
        kakikudashi.innerHTML = kakikudashibun.text 

    } catch (error) {
        console.error('Erorr:', error);
            alert("Failure; check the console");
    }
};
