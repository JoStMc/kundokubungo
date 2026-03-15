const rawInput = document.getElementById('input');
const inputButton = document.getElementById('btn-input');
const tategaki = document.querySelector('.tategaki');

async function renderTategaki(sentence) {
    const container = document.getElementById("tategaki");
    container.innerHTML = ""; 

    characters = sentence.Characters
    characters.forEach((char, index) => {
        const box = document.createElement("div");
        box.className = "kanji-box";
        box.innerHTML = `
<span>${char.Kanji}</span>
<textarea class="kaeriten"></textarea>
<!-- <textarea class="okurigana2"></textarea> -->
<textarea class="okurigana"></textarea>
`;
        box.dataset.index = index; 
        container.appendChild(box);
    });}

inputButton.addEventListener('click', async () => {
    const text = rawInput.value;
    if (!text) return; 

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
