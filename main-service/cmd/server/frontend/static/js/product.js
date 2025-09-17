
/** Product page UI only: gallery switching, add-to-cart button toast */
(function(){
  const mainImg = document.querySelector('[data-main-img]');
  document.querySelectorAll('[data-thumb]').forEach(thumb => {
    thumb.addEventListener('click', () => {
      if(mainImg) mainImg.src = thumb.src;
    });
  });

  const qtyInput = document.querySelector('input[name="qty"]');
  document.querySelectorAll('[data-qty]').forEach(btn => {
    btn.addEventListener('click', () => {
      const dir = btn.dataset.qty;
      const cur = parseInt(qtyInput.value || '1', 10);
      qtyInput.value = Math.max(1, cur + (dir === 'inc' ? 1 : -1));
    });
  });
})();
