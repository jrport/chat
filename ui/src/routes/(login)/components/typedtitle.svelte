<script>
  export let text = "Texto padrao";
  export let typingSpeed = 50;

  let displayedText = "";
  let currentIndex = 0;

  function typeText() {
    if (currentIndex < text.length) {
      displayedText += text[currentIndex];
      currentIndex++;
      setTimeout(typeText, typingSpeed);
    }
  }

    typeText();
</script>

<div>
  {displayedText}
</div>

