
/** Catalog page: only UI interactions. No data mocking. */
(function(){
  const params = new URLSearchParams(location.search);
  const q = params.get('q');
  if(q){
    const el = document.querySelector('[data-query]');
    if(el) el.textContent = q;
  }

  // Filter chips UI
  document.querySelectorAll('.chip').forEach(ch => {
    ch.addEventListener('click', () => ch.classList.toggle('active'));
  });

  // Sort change = soft UI only
  const sort = document.querySelector('[name="sort"]');
  if(sort){
    sort.addEventListener('change', () => {
      // In real app, re-request from backend
      window.showToast('Сортировка применена (требуется бэкенд)');
    });
  }

  // Pagination controls (UI only)
  document.querySelectorAll('[data-page]').forEach(btn => {
    btn.addEventListener('click', (e) => {
      e.preventDefault();
      window.showToast('Пагинация требует бэкенд');
    });
  });
})();
