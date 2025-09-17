
/** Common UI logic (no data mocking) */
(function(){
  const body = document.body;

  // Mobile menu (placeholder)
  const menuBtn = document.querySelector('[data-menu-toggle]');
  if(menuBtn){
    menuBtn.addEventListener('click', () => {
      document.documentElement.classList.toggle('menu-open');
    });
  }

  // Toast helper
  window.showToast = (msg, timeout=2000) => {
    let t = document.querySelector('.toast');
    if(!t){
      t = document.createElement('div');
      t.className = 'toast';
      document.body.appendChild(t);
    }
    t.textContent = msg;
    t.classList.add('show');
    setTimeout(() => t.classList.remove('show'), timeout);
  };

  // Intercept any button with data-requires-backend to show a toast
  document.addEventListener('click', (e) => {
    const btn = e.target.closest('[data-requires-backend]');
    if(btn){
      e.preventDefault();
      showToast('Эта функция требует готового бэкенда');
    }
  });

  // Simple "active" state toggles for favorites/compare on UI only
  document.addEventListener('click', (e) => {
    const fav = e.target.closest('[data-fav-toggle]');
    if(fav){
      e.preventDefault();
      fav.classList.toggle('active');
      showToast(fav.classList.contains('active') ? 'Добавлено в избранное (UI)' : 'Удалено из избранного (UI)');
    }
    const cmp = e.target.closest('[data-compare-toggle]');
    if(cmp){
      e.preventDefault();
      cmp.classList.toggle('active');
      showToast(cmp.classList.contains('active') ? 'Добавлено к сравнению (UI)' : 'Удалено из сравнения (UI)');
    }
  });

  // Search form placeholder
  const searchForm = document.querySelector('[data-search-form]');
  if(searchForm){
    searchForm.addEventListener('submit', (e) => {
      e.preventDefault();
      const q = searchForm.querySelector('input[name="q"]').value.trim();
      if(q){
        // Navigate to catalog with query param (real search should call backend)
        window.location.href = 'catalog.html?q=' + encodeURIComponent(q);
      }
    });
  }

})();
