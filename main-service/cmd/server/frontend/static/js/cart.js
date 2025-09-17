
/** Cart page — UI only. Shows stubs and disables checkout until backend is ready. */
(function(){
  // Quantity buttons
  document.querySelectorAll('[data-cart-qty]').forEach(btn => {
    btn.addEventListener('click', () => {
      const row = btn.closest('[data-row]');
      const input = row.querySelector('input[type="number"]');
      let v = parseInt(input.value || '1', 10);
      if(btn.dataset.cartQty === 'inc') v++;
      else v = Math.max(1, v-1);
      input.value = v;
      window.showToast('Количество изменено (требуется бэкенд для сохранения)');
    });
  });
})();
